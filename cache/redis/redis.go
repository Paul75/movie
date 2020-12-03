package redis

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisDB struct {
	client *redis.Client
	exp    time.Duration
}

var ctx = context.Background()

// NewRedisDB return a RedisDB from the given params DNS and exp.
// exp represents some seconds.
func NewRedisDB(dns string, exp int) *RedisDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     dns,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Println("ping", rdb.Ping(ctx))
	log.Println("flush all data from redis.", rdb.FlushAll(ctx))

	return &RedisDB{
		client: rdb,
		exp:    time.Duration(exp) * time.Second,
	}
}

func (r *RedisDB) Set(key, val string) error {
	return r.client.Set(ctx, key, val, r.exp).Err()
}

func (r *RedisDB) Get(key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
