package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	workerCount := 5

	for i := 0; i < workerCount; i++ {
		go worker(i, data)
	}

	for i := 0; i < 5; i++ {
		data <- i
	}
	close(data)
}

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
