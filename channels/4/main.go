package main

import "fmt"

func main() {
	channel := make(chan int)
	// channel <- 10 - Deadlock because it needs to run on go routine
	go func() {
		channel <- 10
	}()
	fmt.Println(<-channel)
}
