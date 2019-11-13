package redis

import (
	"fmt"
	"sync"
	"time"

	redis "github.com/go-redis/redis"
)

//Config redis的配置文件
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Auth     string `json:"auth"`
	Db       int    `json:"db"`
	PoolSize int    `json:"poolSize"`
}

var config Config

//InitRedis 设置Redis配置文件
func InitRedis(conf Config) {
	config = conf
}

var redisOnce sync.Once
var client *redis.Client

//ErrKeyNotExists not exists
var ErrKeyNotExists = redis.Nil

//GetRedisClient get the client of redis
func getRedisClient() *redis.Client {
	redisOnce.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:       fmt.Sprintf("%s:%d", config.Host, config.Port),
			Password:   config.Auth,
			DB:         config.Db,
			MaxRetries: 2,
			PoolSize:   config.PoolSize,
		})
	})
	return client
}

//Get get redis value
func Get(key string) (string, error) {
	return getRedisClient().Get(key).Result()
}

//Set set redis key value
func Set(key string, val string, expiration time.Duration) error {
	return getRedisClient().Set(key, val, expiration).Err()
}

//Del delete the key
func Del(key string) error {
	return getRedisClient().Del(key).Err()
}

//TTL change the Ttl
func TTL(key string) (time.Duration, error) {
	r := getRedisClient().TTL(key)
	return r.Val(), r.Err()
}

//Client return the raw redis client
func Client() *redis.Client {
	return getRedisClient()
}

func RPush(key string, value string) error {

	return getRedisClient().RPush(key, value).Err()
}

func LPush(key string, value string) error {
	return getRedisClient().LPush(key, value).Err()
}

func RPop(key string) error {
	return getRedisClient().RPop(key).Err()
}

func LPop(key string) error {
	return getRedisClient().LPop(key).Err()
}

func LLen(key string) (int64, error) {
	r := getRedisClient().LLen(key)
	return r.Val(), r.Err()
}

func LRem(key string, count int64, value string) error {
	return getRedisClient().LRem(key, count, value).Err()
}
