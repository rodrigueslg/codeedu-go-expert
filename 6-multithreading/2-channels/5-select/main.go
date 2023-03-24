package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID      int64
	Message string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1) // atomic add (mutex)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1) // atomic add (mutex)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case m := <-c1: // RabbitMQ
			fmt.Printf("Received from RabbitMQ | ID: %d: - Msg %s\n", m.ID, m.Message)

		case m := <-c2: // Kafka
			fmt.Printf("Received from Kafka | ID: %d: - Msg %s\n", m.ID, m.Message)

		case <-time.After(time.Second * 5):
			println("timeout")
		}
	}
}
