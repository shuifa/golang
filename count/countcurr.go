package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func atoAdd(counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(counter, 1)
	}
}

func mutAdd(counter *int64, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mtx.Lock()
		*counter++
		mtx.Unlock()
	}
}

func doMuxAdd() int64 {
	var counter int64 = 0
	var wg = sync.WaitGroup{}
	var mtx = sync.Mutex{}
	wg.Add(2)
	go mutAdd(&counter, &wg, &mtx)
	go mutAdd(&counter, &wg, &mtx)
	wg.Wait()
	return counter
}

func doAtoAdd() int64 {
	var counter int64 = 0
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go atoAdd(&counter, &wg)
	go atoAdd(&counter, &wg)
	wg.Wait()
	return counter
}

func main() {
	fmt.Println(doMuxAdd())
	fmt.Println(doAtoAdd())
}
