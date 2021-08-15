package geecache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name      string
	getter    Getter
	mainCache cache
	Peers     PeerPicker
}

var (
	mtx    sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, mainBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("getter is nil")
	}

	mtx.Lock()
	defer mtx.Unlock()

	g := &Group{name, getter, cache{cacheBytes: mainBytes}, nil}

	groups[name] = g

	return g
}

func GetGroup(name string) *Group {
	mtx.RLock()
	defer mtx.RUnlock()
	return groups[name]
}

func (g *Group) Get(key string) (ByteView, error) {

	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if value, ok := g.mainCache.Get(key); ok {
		fmt.Println("从本地获取到缓存啦", g.mainCache.lru)
		return value.(ByteView), nil
	}

	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	if g.Peers != nil {
		if peer, ok := g.Peers.PickPeer(key); ok {

			if bytes, err := peer.Get(g.name, key); err == nil {
				return ByteView{b: bytes}, nil
			} else {
				log.Printf("[GeeCache] failed to load value from peer=%v, err=%v \n", g.Peers, err)
			}
		}
	}
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (value ByteView, err error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}

	value = ByteView{b: cloneBytes(bytes)}

	g.populateCache(key, value)

	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.Add(key, value)
}

func (g *Group) RegisterPeers(peers PeerPicker) {
	if peers == nil {
		panic("peerPicker should not be nil")
	}
	g.Peers = peers
}
