package geecache

import (
	"fmt"
	"log"
	"sync"

	pb "github.com/oushuifa/golang/better/cache/geecachepb"
	"github.com/oushuifa/golang/better/cache/singleflight"
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
	loader    *singleflight.Group
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

	g := &Group{name, getter, cache{cacheBytes: mainBytes}, nil, &singleflight.Group{}}

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

	view, er := g.loader.Do(key, func() (interface{}, error) {

		if g.Peers != nil {
			if peer, ok := g.Peers.PickPeer(key); ok {
				req := &pb.Request{
					Group: g.name,
					Key:   key,
				}
				res := &pb.Response {}
				if err := peer.Get(req, res); err == nil {
					return ByteView{b: res.Value}, nil
				} else {
					log.Printf("[GeeCache] failed to load value from peer=%v, err=%v \n", g.Peers, err)
				}
			}
		}

		return g.getLocally(key)
	})

	if er == nil {
	    return view.(ByteView), nil
	}

	return ByteView{}, nil

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
