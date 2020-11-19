package redis

import (
	"fmt"
	"ginn/config"
	"github.com/go-redis/redis"
)

var cache *redis.Client

func init() {
	cache = NewRedis()
}

func NewRedis() (cache *redis.Client) {
	redisCfg := config.Gconfig.Redis
	cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisCfg.Host, redisCfg.Port),
		Password: redisCfg.PassWord,
		DB:       redisCfg.DbName,
		PoolSize: redisCfg.PoolSize,
	})
	_, err := cache.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis ping failed! error:%v", err))
	}
	return
}

func CloseRedis() {
	_ = cache.Close()
}
