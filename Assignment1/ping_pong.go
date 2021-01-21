package main

import (
	"fmt"
	"time"
)

// The pinger prints a ping & waits for a pong
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

// The ponger prints a pong & waits fot a ping
func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func main() {
	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	// The main goroutine starts the ping/pong by sending into the ping channel - The ball starts from Ping
	ping <- 1

	for {
		// Block th emain thread until an interrupt
		time.Sleep(time.Second)
	}
}
