package goo_sms

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var __cache *cache

func InitCache(client *redis.Client) {
	__cache = &cache{
		redis: client,
	}
}

type cache struct {
	redis *redis.Client
}

func (ca *cache) set(appid, mobile, action, code string, expireIn int64) error {
	key := fmt.Sprintf("%s_%s_%s", appid, mobile, action)
	return ca.redis.Set(key, code, time.Duration(expireIn)*time.Second).Err()
}

func (ca *cache) get(appid, mobile, action string) string {
	key := fmt.Sprintf("%s_%s_%s", appid, mobile, action)
	return ca.redis.Get(key).Val()
}

func (ca *cache) del(appid, mobile, action string) error {
	key := fmt.Sprintf("%s_%s_%s", appid, mobile, action)
	return ca.redis.Del(key).Err()
}
