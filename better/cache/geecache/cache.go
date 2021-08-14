package geecache

import (
	"sync"

	"github.com/oushuifa/golang/better/cache/lru"
)

type cache struct {
	mtx        sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) Add(key string, value lru.Value) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}

	c.lru.Add(key, value)
}

func(c * cache) Get(key string) (value lru.Value, ok bool){

    c.mtx.Lock()

    defer c.mtx.Unlock()

    if c.lru == nil {
        return
    }

    return c.lru.Get(key)
}