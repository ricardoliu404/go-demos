package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64
	wg  sync.WaitGroup
	//互斥锁
	mutex sync.Mutex
)

func addCount() {
	defer wg.Done()
	for i := 0 ; i < 2 ; i++ {
		// 加锁
		mutex.Lock()
		{
			// 安全的非原子操作
			value := counter
			runtime.Gosched()
			value ++
			counter = value
		}
		//释放锁
		mutex.Unlock()
	}
}

func main() {
	wg.Add(2)
	go addCount()
	go addCount()
	wg.Wait()
	fmt.Printf("counter: %d\n", counter)
}