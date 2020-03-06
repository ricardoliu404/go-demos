package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}
		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s missed.", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball ++
		court <- ball
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	court := make(chan int)
	wg.Add(2)
	go player("liuzhx", court)
	go player("chail", court)
	court <- 1
	wg.Wait()
}
