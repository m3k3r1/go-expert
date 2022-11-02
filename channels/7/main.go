package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel1 := generateMsg("P1")
	channel2 := generateMsg("P2")
	channel := funnel(channel1, channel2)
	for i := 0; i < 20; i++ {
		fmt.Println(<-channel)
	}
	fmt.Println("Finished")
}

func generateMsg(s string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- fmt.Sprintf("String -> %s  Value -> %d", s, i)
			time.Sleep(time.Duration(rand.Intn(255)) * time.Millisecond)
		}
	}()
	return channel
}

func funnel(channel1, channel2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for {
			channel <- <-channel1
		}
	}()
	go func() {
		for {
			channel <- <-channel2
		}
	}()

	return channel
}
