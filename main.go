package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shuifa/golang/pool"
)

var (
	counter int32
	wg sync.WaitGroup
)

type DBConnection struct {
	id int32
}

func (c DBConnection) Close() error {
	fmt.Println("数据库连接关闭啦")
	return nil
}

func Factory() (io.Closer, error) {
	atomic.AddInt32(&counter, 1)
	return &DBConnection{id: counter}, nil
}

func performQuery(query int, pool *pool.Pool)  {
	defer wg.Done()
	resources, err := pool.AcquireResources()
	if err != nil {
		log.Fatalln("获取资源失败", err)
		return
	}
	defer pool.ReleaseResource(resources)

	time.Sleep(time.Duration(rand.Int() % 2) * time.Second)

	fmt.Println("查询完成 id" + fmt.Sprint(query))
}

func main() {
	p, err := pool.New(Factory, 5)
	if err != nil {
		log.Fatalln("创建资源失败", err)
	}
	wg.Add(15)
	for id := 0; id < 15; id++ {
		go performQuery(id, p)
	}
	wg.Wait()
}

func creatTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Println("task complete # ", id)
	}
}