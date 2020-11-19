package redis

import "time"

func SetKey(key string, value interface{}, expiration time.Duration) (err error) {
	err = cache.Set(key, value, expiration).Err()
	return
}

func GetKey(key string) (value string, err error) {
	value, err = cache.Get(key).Result()
	return
}

func DelKey(key ...string) (err error) {
	err = cache.Del(key...).Err()
	return
}
