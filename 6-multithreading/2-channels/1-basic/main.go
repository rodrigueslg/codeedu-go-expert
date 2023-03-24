package main

import "time"

// Thread 1
func main() {

	// Thread 2
	c := make(chan string) // new emtpy channel

	go func() {
		c <- "Hello World" // fill channel
	}()

	msg := <-c // read channel
	println(msg)

	time.Sleep(10 * time.Second)
}
