package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ping(ch chan string, ack chan string, iterations int) {
	for i := 0; i < iterations; i++ {
		fmt.Println("ping")
		ch <- "pong" // Send signal to pong

		// Wait for acknowledgment from pong before continuing
		<-ack

		// Random wait time between 100ms and 1000ms
		randomWait := time.Duration(rand.Intn(900)+100) * time.Millisecond
		time.Sleep(randomWait)
	}
	ch <- "done" // Signal pong to stop
}

func pong(ch chan string, ack chan string) {
	for {
		msg := <-ch // Wait for message from ping
		if msg == "done" {
			fmt.Println("Received done signal, exiting...")
			return
		}

		// Random wait time between 100ms and 1000ms before responding
		randomWait := time.Duration(rand.Intn(900)+100) * time.Millisecond
		time.Sleep(randomWait)

		fmt.Println(msg)

		// Send acknowledgment back to ping
		ack <- "ready"
	}
}

func main() {
	ch := make(chan string)
	ack := make(chan string)       // Acknowledgment channel for ping to wait for pong
	iterations := rand.Intn(9) + 1 // Arbitrary number of iterations

	fmt.Printf("Starting ping-pong with %d iterations\n", iterations)

	// Start both goroutines
	go ping(ch, ack, iterations)
	pong(ch, ack)

}
