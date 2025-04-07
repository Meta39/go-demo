package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"redis/distributed"
	"time"
)

/*
redis操作
1.下载go-redis库：go get github.com/redis/go-redis/v9
*/
func main() {
	// 只需在应用初始化时调用一次
	distributed.InitRedisClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 调用 DoWithLockDefault 分布式锁，锁超时时间默认 30 秒，超时后看门狗自动续期
	if err := distributed.DoWithLockDefault("lockName", func() error {
		log.Println("执行业务逻辑...")
		// 模拟业务耗时
		time.Sleep(45 * time.Second)
		log.Println("业务逻辑执行完毕")
		return nil
	}); err != nil {
		log.Printf("业务执行失败：%v", err)
		return
	}

	// 调用 DoWithLockTxDefaultTimeoutAndRetries 包装事务业务逻辑，锁超时时间默认 300 秒，超时后看门狗自动续期，重试次数默认 3 次。
	if err := distributed.DoWithLockTxDefaultTimeoutAndRetries("doWithLockTx", func(pipe redis.Pipeliner) error {
		log.Println("执行事务业务逻辑...")
		// 模拟业务耗时
		time.Sleep(45 * time.Second)
		// 例如：在事务中设置一个键值
		pipe.Set(context.Background(), "someKey", "someValue", 0)
		log.Println("事务业务逻辑执行完毕")
		// 可以在这里继续添加其他 Redis 操作
		return nil
	}); err != nil {
		log.Printf("事务业务执行失败: %v", err)
	}

}
