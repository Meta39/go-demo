package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"redis/config"
	"redis/distributed"
	"redis/redisutil"
	"strconv"
	"time"
)

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
	doWithLockDefault()

	// 调用 DoWithLockTxDefaultTimeoutAndRetries 包装事务业务逻辑，锁超时时间默认 300 秒，超时后不会自动续期，重试次数默认 3 次。【用于多个redis命令原子操作】
	doWithLockTxDefaultTimeoutAndRetries()

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
	if err := distributed.DoWithLockDefault("lockName", func() error {
		log.Println("执行业务逻辑...")
		// 模拟业务耗时
		time.Sleep(3 * time.Second)
		log.Println("业务逻辑执行完毕")
		return nil
	}); err != nil {
		log.Printf("业务执行失败：%v", err)
		return
	}
}

// 调用 DoWithLockTxDefaultTimeoutAndRetries 包装事务业务逻辑，锁超时时间默认 300 秒，超时后不会自动续期，重试次数默认 3 次。【用于多个redis命令原子操作】
func doWithLockTxDefaultTimeoutAndRetries() {
	if err := distributed.DoWithLockTxDefaultTimeoutAndRetries("doWithLockTx", func(pipe redis.Pipeliner) error {
		log.Println("执行事务业务逻辑...")
		// 模拟业务耗时
		time.Sleep(3 * time.Second)
		// 例如：在事务中设置一个键值
		pipe.Set(context.Background(), "someKey", "someValue", 30*time.Second)
		log.Println("事务业务逻辑执行完毕")
		// 可以在这里继续添加其他 Redis 操作
		return nil
	}); err != nil {
		log.Printf("事务业务执行失败: %v", err)
	}
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
