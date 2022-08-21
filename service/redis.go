package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

var ctx context.Context
var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       8,  // use default DB
	})
	ctx = context.Background()
}

// SAdd 向集合添加一个成员
func SAdd(key string, value string) {
	_, err := rdb.SAdd(ctx, key, value).Result()
	if err != nil {
		log.Fatal("sadd error", err)
	}
}

// SRemove 移除集合中一个或多个成员
func SRemove(key string, value string) {
	result, err := rdb.SRem(ctx, key, value).Result()
	if err != nil {
		fmt.Println("srem error")
	}
	fmt.Println("srem success", result)
}

// IsMember 判断 member 元素是否是集合 key 的成员
func IsMember(key string, id string) bool {
	result := rdb.SIsMember(ctx, key, id).Val()
	return result
}

// HDecr 递减
func HDecr(key string, value string) {
	_, err := rdb.HIncrBy(ctx, key, value, -1).Result()
	if err != nil {
		log.Fatal("HDecr error")
	}
}

// HIncr 递增
func HIncr(key string, value string) {
	_, err := rdb.HIncrBy(ctx, key, value, 1).Result()
	if err != nil {
		log.Fatal("HDecr error")
	}
}
