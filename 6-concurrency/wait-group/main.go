package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// thread 1 (main is an implicit thread)

	wg := &sync.WaitGroup{}
	wg.Add(25) // 25 credits (A: 10, B: 10, anonymous: 5)

	// thread 2
	go Task("A", wg)

	// thread 3
	go Task("A", wg)

	// thread 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(time.Second)
			wg.Done()
		}
	}()

	wg.Wait()
}

func Task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second)
		wg.Done()
	}
}
