package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg  sync.WaitGroup
)

func addCount() {
	defer wg.Done()
	for i := 0 ; i < 2 ; i++ {
		//          原子操作
		// 原子自增
		atomic.AddInt64(&counter, 1)
		// 让出时间片
		runtime.Gosched()
	}
}

func main() {
	wg.Add(2)
	go addCount()
	go addCount()
	wg.Wait()
	fmt.Printf("counter: %d\n", counter)
}
