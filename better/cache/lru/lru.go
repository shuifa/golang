package lru

import (
    "container/list"
)

type Cache struct {
    maxBytes  int64
    nbytes    int64
    ll        *list.List
    cache     map[string]*list.Element
    OnEvicted func(key string, value Value)
}

type entry struct {
    key   string
    value Value
}

type Value interface {
    Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
    return &Cache{
        maxBytes:  maxBytes,
        nbytes:    0,
        ll:        list.New(),
        cache:     make(map[string]*list.Element),
        OnEvicted: onEvicted,
    }
}

func (c *Cache) Get(key string) (value Value, ok bool) {
    if ele, ok := c.cache[key]; ok {
        c.ll.MoveToFront(ele)
        return ele.Value.(*entry).value, ok
    }
    return nil, false
}

func (c *Cache) removeOldest() {
    if ele := c.ll.Back(); ele != nil {
        c.ll.Remove(ele)
        entry := ele.Value.(*entry)
        delete(c.cache, entry.key)

        c.nbytes -= int64(len(entry.key)) + int64(entry.value.Len())
        if c.OnEvicted != nil {
            c.OnEvicted(entry.key, entry.value)
        }
    }
}

func(c * Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		// 更新
		c.ll.MoveToFront(ele)

		entry := ele.Value.(* entry)
		c.nbytes += int64(value.Len()) - int64(entry.value.Len())
		entry.value = value

	} else {
	    // 插入
		ele := c.ll.PushFront(&entry {
		    key : key,
		    value : value,
		})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	if c.maxBytes != 0 &&  c.nbytes > c.maxBytes {
		c.removeOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}

