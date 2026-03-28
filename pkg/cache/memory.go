package cache

import (
    "sync"
    "time"
)

type memoryCache struct {
    mu    sync.RWMutex
    items map[string]cacheItem
}

type cacheItem struct {
    value     []byte
    expiresAt time.Time
}

func NewMemoryCache() Cache {
    return &memoryCache{
        items: make(map[string]cacheItem),
    }
}

func (c *memoryCache) Get(key string) ([]byte, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    item, exists := c.items[key]
    if !exists {
        return nil, false
    }

    if time.Now().After(item.expiresAt) {
        return nil, false
    }

    return item.value, true
}

func (c *memoryCache) Set(key string, value []byte, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.items[key] = cacheItem{
        value:     value,
        expiresAt: time.Now().Add(ttl),
    }
}

func (c *memoryCache) Delete(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()

    delete(c.items, key)
}
