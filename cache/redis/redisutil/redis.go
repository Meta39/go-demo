package redisutil

import (
	"context"
	"github.com/redis/go-redis/v9"
	"redis/config"
	"time"
)

//redis工具类，缺少的函数在这添加，不要单独操作。

// ---------------------- 字符串（String）操作 ----------------------

// Set 设置字符串键值，过期时间支持0表示永不过期（原子操作）
func Set(key string, value interface{}, expiration time.Duration) error {
	return config.RedisClient.Set(context.Background(), key, value, expiration).Err()
}

// Get 获取字符串键值（原子操作）
func Get(key string) (string, error) {
	return config.RedisClient.Get(context.Background(), key).Result()
}

func GetByte(key string) ([]byte, error) {
	return config.RedisClient.Get(context.Background(), key).Bytes()
}

// MSet 批量设置多个字符串键值（原子操作，单命令执行）
func MSet(kv map[string]interface{}) error {
	return config.RedisClient.MSet(context.Background(), kv).Err()
}

// MGet 批量获取多个字符串键值（原子操作，单命令执行）
func MGet(keys ...string) ([]interface{}, error) {
	return config.RedisClient.MGet(context.Background(), keys...).Result()
}

// Incr 原子递增计数器（原子操作）
func Incr(key string) (int64, error) {
	return config.RedisClient.Incr(context.Background(), key).Result()
}

// ---------------------- 哈希（Hash）操作 ----------------------

// HSet 设置哈希表单个字段（原子操作）
func HSet(key string, field string, value interface{}) error {
	return config.RedisClient.HSet(context.Background(), key, field, value).Err()
}

// HGet 获取哈希表单个字段（原子操作）
func HGet(key string, field string) (string, error) {
	return config.RedisClient.HGet(context.Background(), key, field).Result()
}

// HMSet 批量设置哈希表多个字段（原子操作，单命令执行）
func HMSet(key string, fields map[string]interface{}) error {
	return config.RedisClient.HMSet(context.Background(), key, fields).Err()
}

// HMGet 批量获取哈希表多个字段（原子操作，单命令执行）
func HMGet(key string, fields ...string) ([]interface{}, error) {
	return config.RedisClient.HMGet(context.Background(), key, fields...).Result()
}

// ---------------------- 列表（List）操作 ----------------------

// LPush 向列表左端插入一个或多个元素（原子操作，单命令执行）
func LPush(key string, values ...interface{}) error {
	return config.RedisClient.LPush(context.Background(), key, values...).Err()
}

// LRange 获取列表指定范围的元素（原子操作）
func LRange(key string, start, stop int64) ([]string, error) {
	return config.RedisClient.LRange(context.Background(), key, start, stop).Result()
}

// ---------------------- 集合（Set）操作 ----------------------

// SAdd 向集合添加一个或多个成员（原子操作，单命令执行）
func SAdd(key string, members ...interface{}) error {
	return config.RedisClient.SAdd(context.Background(), key, members...).Err()
}

// SMembers 获取集合所有成员（原子操作）
func SMembers(key string) ([]string, error) {
	return config.RedisClient.SMembers(context.Background(), key).Result()
}

// ---------------------- 有序集合（ZSet）操作 ----------------------

// ZAdd 向有序集合添加一个或多个成员（原子操作，单命令执行）
func ZAdd(key string, members ...redis.Z) error {
	return config.RedisClient.ZAdd(context.Background(), key, members...).Err()
}

// ZRangeByScore 按分数范围获取有序集合成员（原子操作）
func ZRangeByScore(key string, min, max string) ([]string, error) {
	return config.RedisClient.ZRangeByScore(context.Background(), key, &redis.ZRangeBy{Min: min, Max: max}).Result()
}

// ---------------------- 事务（Transaction） ----------------------

// TxPipelined 开启事务，返回的事务对象在调用Exec()时原子执行
// 原子性说明：事务内所有命令在Exec()调用时原子执行
func TxPipelined(fn func(redis.Pipeliner) error) error {
	_, err := config.RedisClient.TxPipelined(context.Background(), fn)
	return err
}

// ---------------------- Lua脚本 ----------------------

// Eval 执行Lua脚本（原子操作，脚本整体原子执行）
func Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	return config.RedisClient.Eval(context.Background(), script, keys, args...).Result()
}

// ---------------------- 管道（Pipeline） ----------------------

// Pipelined 开启管道（非原子操作，用于批量命令发送）
// 原子性说明：管道中的命令独立执行，不保证原子性
func Pipelined(fn func(redis.Pipeliner) error) error {
	_, err := config.RedisClient.Pipelined(context.Background(), fn)
	return err
}

// ---------------------- 键管理 ----------------------

// Del 删除一个或多个键（原子操作，单命令执行）
func Del(keys ...string) error {
	return config.RedisClient.Del(context.Background(), keys...).Err()
}

// Expire 设置键的过期时间（原子操作）
func Expire(key string, expiration time.Duration) error {
	return config.RedisClient.Expire(context.Background(), key, expiration).Err()
}

// ---------------------- Pub/Sub ----------------------

// Publish 向频道发布消息（非数据操作，无原子性要求）
func Publish(channel string, message interface{}) error {
	return config.RedisClient.Publish(context.Background(), channel, message).Err()
}

// Subscribe 订阅一个或多个频道（非数据操作，无原子性要求）
func Subscribe(channels ...string) *redis.PubSub {
	return config.RedisClient.Subscribe(context.Background(), channels...)
}
