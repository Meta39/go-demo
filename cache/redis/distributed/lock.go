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

// DoWithLockTxDefaultTimeout 默认超时时间，重试次数自定义【用于多个redis命令原子操作】
func DoWithLockTxDefaultTimeout(lockKey string, retries int, businessLogic func(pipe redis.Pipeliner) error) error {
	return DoWithLockTx(lockKey, defaultTxExpiration, retries, businessLogic)
}

// DoWithLockTxDefaultRetries 默认重试次数，超时时间自定义【用于多个redis命令原子操作】
func DoWithLockTxDefaultRetries(lockKey string, expiration time.Duration, businessLogic func(pipe redis.Pipeliner) error) error {
	return DoWithLockTx(lockKey, expiration, defaultRetries, businessLogic)
}

// DoWithLockTxDefaultTimeoutAndRetries 默认超时时间、重试次数。【用于多个redis命令原子操作】
func DoWithLockTxDefaultTimeoutAndRetries(lockKey string, businessLogic func(pipe redis.Pipeliner) error) error {
	return DoWithLockTx(lockKey, defaultTxExpiration, defaultRetries, businessLogic)
}

// DoWithLockTx 包装事务业务逻辑执行，加锁、事务控制及释放锁（无看门狗续约，因为WATCH机制不允许，因此需要合理的设置超时时间）【用于多个redis命令原子操作】
func DoWithLockTx(lockKey string, expiration time.Duration, retries int, businessLogic func(pipe redis.Pipeliner) error) error {
	if retries <= 0 {
		retries = defaultRetries
	}

	var lastErr error
	for i := 0; i < retries; i++ {
		lock := NewDistributedLock(lockKey, expiration)
		//尝试加锁（无自动续约，需要合理的设置超时时间）
		locked, err := lock.LockWithoutAutoRenew()
		if err != nil {
			lastErr = fmt.Errorf("lock error: %w", err)
			continue
		}
		if !locked {
			//没有获取到锁跳过本次循环，或抛出错误。（推荐跳过本次循环）
			//lastErr = errors.New("unable to acquire lock")
			continue
		}

		// 开启事务监控锁 key
		err = lock.client.Watch(lock.ctx, func(tx *redis.Tx) error {
			_, err := tx.TxPipelined(context.Background(), func(pipe redis.Pipeliner) error {
				return businessLogic(pipe)
			})
			return err
		}, lock.key)

		// 主动释放锁（如果还未自动过期）
		if uErr := lock.Unlock(); uErr != nil {
			log.Printf("unlock error: %v", uErr)
		}

		if err == nil {
			return nil
		}

		lastErr = err
		log.Printf("事务执行失败，第 %d 次重试，error: %v", i+1, err)
		//停顿时间（推荐不设置，因为集群模式下，大家都会争夺资源，所以不用控制停顿时间。）
		//time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("transaction failed after %d retries, last error: %w", retries, lastErr)
}

// LockWithoutAutoRenew 尝试加锁（无自动续约）
func (l *DistributedLock) LockWithoutAutoRenew() (bool, error) {
	ok, err := l.client.SetNX(l.ctx, l.key, l.value, l.expiration).Result()
	if err != nil {
		return false, err
	}
	return ok, nil
}
