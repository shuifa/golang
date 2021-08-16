package singleflight

import (
	"sync"
)

type call struct {
	wg    sync.WaitGroup
	value interface{}
	err   error
}

type Group struct {
	mtx sync.Mutex
	m   map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mtx.Lock()

	if g.m == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
        g.mtx.Unlock()
        c.wg.Wait()
        return c.value, c.err
	}

	c := new(call)
	c.wg.Add(1)

	g.m[key] = c
	g.mtx.Unlock()

	c.value, c.err = fn()
    c.wg.Done()

	g.mtx.Lock()

	delete(g.m, key)

	g.mtx.Unlock()

    return c.value, c.err
}
