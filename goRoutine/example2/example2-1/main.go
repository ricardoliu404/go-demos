package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64
	wg  sync.WaitGroup
)

func addCount() {
	defer wg.Done()
	for i := 0 ; i < 2 ; i++ {
		//          非原子操作
		// 读数据
		value := counter
		// 让出时间片
		runtime.Gosched()
		// 临时值自增
		value ++
		//回写
		counter = value
	}
}

func main() {
	wg.Add(2)
	go addCount()
	go addCount()
	wg.Wait()
	fmt.Printf("counter: %d\n", counter)
}
