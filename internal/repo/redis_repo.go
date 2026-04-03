package repo

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

// InitRedis 初始化redis
func InitRedis(addr, password string, db int) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// AcquireLock 获取锁
func AcquireLock(key string, ttl time.Duration) (bool, error) {
	ok, err := Rdb.SetNX(Ctx, key, 1, ttl).Result()
	return ok, err
}

func ReleaseLock(key string) error {
	return Rdb.Del(Ctx, key).Err()
}
