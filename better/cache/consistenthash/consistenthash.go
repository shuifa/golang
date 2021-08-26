package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type HashFunc func(data []byte) uint32

type Map struct {
	hashFunc HashFunc
	keys     []int
	hashMap  map[int]string
	replicas int
}

func NewMap(replicas int, hashFunc HashFunc) *Map {

	if hashFunc == nil {
		hashFunc = crc32.ChecksumIEEE
	}

	return &Map{
		hashFunc: hashFunc,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
}

func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hashValue := int(m.hashFunc([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hashValue)
			m.hashMap[hashValue] = key
		}
	}

	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {

	if len(m.keys) == 0 {
		return ""
	}

	hashValue := int(m.hashFunc([]byte(key)))

	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hashValue % m.keys[len(m.keys)-1] })

	return m.hashMap[m.keys[idx%len(m.keys)]]
}
