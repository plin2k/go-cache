package go_cache

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
)

type redisCache struct {
	client *redis.Client
	keys   sync.Map
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
		keys:   sync.Map{},
	}, nil
}

func (r *redisCache) Get(key string) (interface{}, error) {

	if !r.keyExists(key) {
		return nil, errNotFound
	}

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

	r.keyStore(key, duration)

	return nil
}

func (r *redisCache) Delete(key string) error {
	err := r.client.Del(key).Err()

	r.keys.Delete(key)

	return err
}

func (r *redisCache) Flush() error {
	err := r.client.FlushDB().Err()

	r.keys = sync.Map{}

	return err
}

func (r *redisCache) keyExists(key string) bool {

	expiration, exists := r.keys.Load(key)
	if !exists {
		return false
	}

	if time.Now().After(expiration.(time.Time)) {
		r.keys.Delete(key)
		return false
	}

	return true
}

func (r *redisCache) keyStore(key string, duration time.Duration) {
	var nowTime = time.Now()

	if duration == 0 {
		nowTime = nowTime.Add(1 << 16 * time.Hour)
	}

	r.keys.Store(key, nowTime.Add(duration))
}
