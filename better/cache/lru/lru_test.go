package lru

import (
	"log"
	"testing"
)

type String string

func (s String) Len() int {
	return len(s)
}

func TestCache_Get(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("testKey1", String("testValue1"))

	if vaule, ok := lru.Get("testKey1"); !ok || vaule != String("testValue1") {
		t.Fatalf("cache hit key1=1234 failed")
	}

	if _, ok := lru.Get("test123"); ok {
		t.Fatalf("run err , test123 should not be fonud")
	}
}

func TestCache_removeOldest(t *testing.T) {
	k1, k2, k3 := "test1", "test2", "test3"
	v1, v2, v3 := "value1", "value2", "value3"
	lru := New(int64(len(k1)+len(v2)+len(k2)+len(v1)), func(key string, value Value) {
		log.Printf("超过内存被删除了哦 key=%s, vaule=%v", key, value)
	})

	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get(k1); ok {
		t.Fatalf("run err , %s should not be fonud", k1)
	}

}
