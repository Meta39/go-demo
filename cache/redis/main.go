package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"redis/config"
	"redis/distributed"
	"redis/redisutil"
	"strconv"
	"sync"
	"time"
)

const numWorkers = 1000

var wg sync.WaitGroup

/*
redis操作
1.下载go-redis库：go get github.com/redis/go-redis/v9
*/
func main() {
	// 只需在应用初始化时调用一次
	config.InitRedisClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 调用 DoWithLockDefault 分布式锁，锁超时时间默认 30 秒，超时后看门狗自动续期【分布式锁】
	for i := 0; i < numWorkers; i++ { //开启1000个协程，看看分布式锁是否成功。
		wg.Add(1) // 增加计数器
		go doWithLockDefault()
	}
	wg.Wait() // 等待所有worker完成

	//多个redis命令原子操作使用【Watch + 事务管道，使用 GET + SET + WATCH 来实现Key递增效果，类似命令 INCR】
	key := "key"
	_ = increment(key, 3, func(pipe redis.Pipeliner) error {
		pipe.Set(context.Background(), key, 1, 15*time.Second)
		pipe.Get(context.Background(), key)
		//后续还可以有多个命令，但是多个命令都要与当前key有关，因为简体的就是这个key
		log.Println("increment key:", key)
		return nil
	})

	// 原子操作示例（单条命令基本上都是原子操作）
	redisSet()

	// 非原子批量操作（普通管道）【一次性操作多条redis命令时不推荐】
	pipelined()

	// 原子事务操作（事务管道）【一次性操作多条redis命令时推荐】
	txPipelined()

}

type user struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

// 调用 DoWithLockDefault 分布式锁，锁超时时间默认 30 秒，超时后看门狗自动续期
func doWithLockDefault() {
	defer wg.Done() // 完成后通知WaitGroup
	err := distributed.DoWithLockDefault("lockName", func() error {
		log.Println("执行业务逻辑...")
		// 模拟业务耗时
		time.Sleep(3 * time.Second)
		log.Println("业务逻辑执行完毕")
		return nil
	})
	switch {
	case errors.Is(err, distributed.ErrLockNotAcquired):
		log.Println("获取锁失败（未抢到锁）")
		return
	case err != nil:
		log.Printf("业务执行失败：%v", err)
		return
	default:
		return
	}
}

// 使用 GET + SET + WATCH 来实现Key递增效果，类似命令 INCR
func increment(key string, maxRetries int, fn func(pipe redis.Pipeliner) error) error {
	// 事务函数
	txf := func(tx *redis.Tx) error {
		n, err := tx.Get(context.Background(), key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		n++

		//fn为业务逻辑
		_, err = tx.TxPipelined(context.Background(), fn)
		return err
	}

	//maxRetries重试次数
	for i := 0; i < maxRetries; i++ {
		err := config.RedisClient.Watch(context.Background(), txf, key)
		if err == nil {
			// Success.
			return nil
		}
		if errors.Is(err, redis.TxFailedErr) {
			// 乐观锁失败
			continue
		}
		return err
	}

	return errors.New("increment reached maximum number of retries")
}

// 原子操作示例（单条命令基本上都是原子操作）
func redisSet() {
	//1.序列化结构体
	u := &user{}
	u.Id = 1
	u.Name = "Meta"
	u.Age = 18
	u.CreatedAt = time.Now()

	userJSON, err := json.Marshal(u) //序列化
	if err != nil {
		panic(fmt.Sprintf("序列化失败: %v", err))
	}
	//2.存储到Redis（设置过期时间30秒）
	setKey := "u:" + strconv.Itoa(u.Id)
	if err := redisutil.Set(setKey, userJSON, 30*time.Second); err != nil {
		log.Printf("redisutil.Set：%v", err)
		return
	}
	//3.从Redis读取数据
	v, _ := redisutil.GetByte(setKey) //默认的Get是获取string
	log.Printf("获取%v的value:%s\n", setKey, v)
	//4.反序列化JSON到结构体
	var retrievedUser user
	err = json.Unmarshal(v, &retrievedUser)
	if err != nil {
		panic(fmt.Sprintf("反序列化失败: %v", err))
	}

	// 5.打印结果
	log.Printf("从Redis读取的用户数据: %+v\n", retrievedUser)
}

// 非原子批量操作（普通管道）【一次性操作多条redis命令时不推荐】
func pipelined() {
	ctx := context.Background()
	if err := redisutil.Pipelined(func(pipeliner redis.Pipeliner) error {
		pipeliner.Set(ctx, "key1", "Meta39", 30*time.Second)
		v, _ := pipeliner.Get(ctx, "key1").Result()
		log.Println("一次性操作多条redis命令 Pipelined 普通管道（非原子操作） key1:", v)
		return nil
	}); err != nil {
		log.Printf("redisutil.Pipelined：%v", err)
		return
	}
}

// 原子事务操作（事务管道）【一次性操作多条redis命令时推荐】
func txPipelined() {
	ctx := context.Background()
	if err := redisutil.TxPipelined(func(pipeliner redis.Pipeliner) error {
		pipeliner.Set(ctx, "key2", "Meta2", 30*time.Second)
		v2, _ := pipeliner.Get(ctx, "key2").Result()
		log.Println("一次性操作多条redis命令 TxPipelined 事务管道（原子操作） key2:", v2)
		return nil
	}); err != nil {
		log.Printf("redisutil.TxPipelined：%v", err)
		return
	}
}
