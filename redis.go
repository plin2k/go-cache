package go_cache

import (
	"github.com/go-redis/redis"
	"time"
)

type redisCache struct {
	client *redis.Client
}

func NewRedis(addr, password string, db int) (*redisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &redisCache{
		client: client,
	}, nil
}

func (r *redisCache) Get(key string) (interface{}, error) {

	val, err := r.client.Get(key).Result()
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (r *redisCache) Set(key string, val interface{}, duration time.Duration) error {
	err := r.client.Set(key, val, duration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisCache) Delete(key string) error {
	err := r.client.Del(key).Err()

	return err
}

func (r *redisCache) Flush() error {
	err := r.client.FlushDB().Err()

	return err
}
