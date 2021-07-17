package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrPoolClosed = errors.New("pool has closed")

type Pool struct {
	Closed bool
	factory func()(io.Closer, error)
	mtx sync.Mutex
	Resources chan io.Closer
}

func New(factory func()(closer io.Closer, err error), size int) (*Pool, error){
	return &Pool{
		Closed: false,
		Resources: make(chan io.Closer,size),
		factory: factory,
	}, nil
}

func (p *Pool) AcquireResources() (io.Closer, error) {
	select {
	case resources, ok := <- p.Resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("获得一个资源链接", resources)
		return resources, nil
	default:
		fmt.Println("用工厂创建了一个资源")
		return p.factory()
	}
}

func (p *Pool) ReleaseResource(resource io.Closer) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	select {
	case p.Resources <- resource:
		fmt.Println("资源丢回池子")
	default:
		fmt.Println("当前资源池子满了-=========")
		_ = resource.Close()
	}

}

func (p *Pool) Close()  {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if p.Closed {
		return
	}

	close(p.Resources)
	for resource := range p.Resources {
		_ = resource.Close()
	}
}