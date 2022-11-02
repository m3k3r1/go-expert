package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	concurrency := 50
	in := make(chan int)
	done := make(chan []byte)

	go func() {
		i := 0
		for {
			in <- i
			i++
		}
	}()

	for x := 0; x < concurrency; x++ {
		go ProcessWorkers(in, x)
	}

	<-done
}

func ProcessWorkers(in chan int, worker int) {
	for x := range in {
		t := time.Duration(rand.Intn(4) * int(time.Second))
		time.Sleep(t)
		fmt.Println("Worker", worker, ":", x)
	}
}
