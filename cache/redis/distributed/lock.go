package distributed

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"redis/config"
	"time"

	"github.com/google/uuid"
)

// 默认超时时间
const defaultExpiration = 30 * time.Second

// 默认事务超时时间
const defaultTxExpiration = 300 * time.Second

// 默认事务重试次数
const defaultRetries = 3

// DistributedLock 封装了分布式锁实现
type DistributedLock struct {
	client     *redis.Client
	key        string
	value      string        // 唯一标识
	expiration time.Duration // 锁过期时间
	cancelFunc context.CancelFunc
	ctx        context.Context
}

// NewDistributedLock 构造锁对象
// 当传入的 expiration 为 0 时，使用默认 30 秒
func NewDistributedLock(key string, expiration time.Duration) *DistributedLock {
	if expiration <= 0 {
		expiration = defaultExpiration
	}
	return &DistributedLock{
		client:     config.RedisClient,
		key:        key,
		value:      uuid.NewString(),
		expiration: expiration,
		ctx:        context.Background(),
	}
}

// Lock 尝试加锁成功则启动看门狗自动续期
func (l *DistributedLock) Lock() (bool, error) {
	ok, err := l.client.SetNX(l.ctx, l.key, l.value, l.expiration).Result()
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}

	// 锁获取成功，启动看门狗协程续约
	renewCtx, cancel := context.WithCancel(l.ctx)
	l.cancelFunc = cancel
	go l.autoRenew(renewCtx, l.expiration/3)
	return true, nil
}

// autoRenew 看门狗定时续约，确保锁在业务逻辑执行期间不失效
func (l *DistributedLock) autoRenew(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Lua 脚本：只有当前持有者才能续期
	renewScript := redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1] then 
			return redis.call("expire", KEYS[1], ARGV[2]) 
		else 
			return 0 
		end
	`)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			seconds := int64(l.expiration.Seconds())
			res, err := renewScript.Run(l.ctx, l.client, []string{l.key}, l.value, seconds).Result()
			if err != nil || res.(int64) != 1 {
				log.Printf("续约失败：err=%v, result=%v", err, res)
				return
			}
		}
	}
}

// Unlock 释放锁，只有锁持有者才能删除锁
func (l *DistributedLock) Unlock() error {
	if l.cancelFunc != nil {
		l.cancelFunc()
	}

	unlockScript := redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1] then 
			return redis.call("del", KEYS[1])
		else 
			return 0 
		end
	`)
	res, err := unlockScript.Run(l.ctx, l.client, []string{l.key}, l.value).Result()
	if err != nil {
		return err
	}
	if res.(int64) != 1 {
		return errors.New("unlock failed: not lock owner")
	}
	return nil
}

// DoWithLockDefault 默认超时时间为30秒的分布式锁【分布式锁】
func DoWithLockDefault(lockKey string, businessLogic func() error) error {
	return DoWithLock(lockKey, defaultExpiration, businessLogic)
}

// DoWithLock 包装业务逻辑执行，内部负责加锁、看门狗续约及释放锁【分布式锁】
// 参数 expiration 为可选，传 0 则使用默认超时 30 秒
func DoWithLock(lockKey string, expiration time.Duration, businessLogic func() error) error {
	lock := NewDistributedLock(lockKey, expiration)
	locked, err := lock.Lock()
	if err != nil {
		return fmt.Errorf("lock error: %w", err)
	}
	if !locked {
		//没有获取到锁直接返回nil，或抛出错误。(推荐返回nil)
		//return errors.New("unable to acquire lock")
		return nil
	}
	defer func() {
		if err := lock.Unlock(); err != nil {
			log.Printf("unlock error: %v", err)
		}
	}()

	return businessLogic()
}
