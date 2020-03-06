package main

// 基本协程及使用

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg sync.WaitGroup
	// CountDownLatch
)

func main() {
	runtime.GOMAXPROCS(1)
	// 设置最多1个协程运行

	wg.Add(2)
	// 初始化CountDownLatch
	fmt.Printf("Begin Goroutines\n")

	go func() {
		defer wg.Done()
		// 协程运行结束后CountDown
		for count := 0 ; count < 3 ; count ++ {
			for char := 'a' ; char <= 'z' ; char ++ {
				if char == 'k'{
					runtime.Gosched()
					// 让出时间片给其他协程使用
				}
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		// 协程运行结束后CountDown
		for count := 0 ; count < 3 ; count ++ {
			for char := 'A' ; char <= 'Z' ; char ++ {
				if char == 'K'{
					runtime.Gosched()
					// 让出时间片给其他协程使用
				}
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Printf("Wait for finishing\n")
	wg.Wait()// 直到CountDownLatch倒数到0再结束
}