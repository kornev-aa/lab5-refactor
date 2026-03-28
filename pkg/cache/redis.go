package cache

import (
    "context"
    "time"
    "github.com/redis/go-redis/v9"
)

type RedisCache struct {
    client *redis.Client
    ctx    context.Context
}

func NewRedisCache(addr string) *RedisCache {
    return &RedisCache{
        client: redis.NewClient(&redis.Options{
            Addr: addr,
        }),
        ctx: context.Background(),
    }
}

func (r *RedisCache) Get(key string) ([]byte, bool) {
    val, err := r.client.Get(r.ctx, key).Bytes()
    if err != nil {
        return nil, false
    }
    return val, true
}

func (r *RedisCache) Set(key string, value []byte, ttl time.Duration) {
    r.client.Set(r.ctx, key, value, ttl)
}

func (r *RedisCache) Delete(key string) {
    r.client.Del(r.ctx, key)
}
