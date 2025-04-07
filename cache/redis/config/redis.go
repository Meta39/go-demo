package config

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

var once sync.Once

// RedisClient 为全局单例
var RedisClient *redis.Client

// InitRedisClient 初始化全局 Redis 客户端，应用启动时只需调用一次
func InitRedisClient(opts *redis.Options) {
	once.Do(func() {
		RedisClient = redis.NewClient(opts)
	})
}
