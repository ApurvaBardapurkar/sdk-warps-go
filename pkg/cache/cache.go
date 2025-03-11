package cache

import (
	"fmt"
	"sync"
	"time"
)

// CacheKey is a type for cache keys
type CacheKey string

// Cache key generators
var CacheKey = struct {
	Warp         func(hash string) CacheKey
	RegistryInfo func(key string) CacheKey
	Brand        func(key string) CacheKey
}{
	Warp: func(hash string) CacheKey {
		return CacheKey(fmt.Sprintf("warp:%s", hash))
	},
	RegistryInfo: func(key string) CacheKey {
		return CacheKey(fmt.Sprintf("registry:%s", key))
	},
	Brand: func(key string) CacheKey {
		return CacheKey(fmt.Sprintf("brand:%s", key))
	},
}

// cacheItem represents an item in the cache
type cacheItem struct {
	value      interface{}
	expiration time.Time
}

// WarpCache is a simple in-memory cache for warps and related data
type WarpCache struct {
	items map[CacheKey]cacheItem
	mutex sync.RWMutex
}

// NewWarpCache creates a new WarpCache
func NewWarpCache() *WarpCache {
	cache := &WarpCache{
		items: make(map[CacheKey]cacheItem),
	}

	// Start the cleanup routine
	go cache.cleanup()

	return cache
}

// Get retrieves a value from the cache
func (c *WarpCache) Get(key CacheKey) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil
	}

	// Check if the item has expired
	if time.Now().After(item.expiration) {
		delete(c.items, key)
		return nil
	}

	return item.value
}

// Set adds a value to the cache with the specified TTL in seconds
func (c *WarpCache) Set(key CacheKey, value interface{}, ttlSeconds int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Calculate expiration time
	expiration := time.Now().Add(time.Duration(ttlSeconds) * time.Second)

	c.items[key] = cacheItem{
		value:      value,
		expiration: expiration,
	}
}

// Delete removes a value from the cache
func (c *WarpCache) Delete(key CacheKey) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)
}

// Clear removes all items from the cache
func (c *WarpCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items = make(map[CacheKey]cacheItem)
}

// cleanup periodically removes expired items from the cache
func (c *WarpCache) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mutex.Lock()

		for key, item := range c.items {
			if time.Now().After(item.expiration) {
				delete(c.items, key)
			}
		}

		c.mutex.Unlock()
	}
} 