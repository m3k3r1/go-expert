package main

import "fmt"

func main() {
	c := generate(4, 6)

	d1 := divide(c)
	d2 := divide(c)

	fmt.Println(<-d1)
	fmt.Println(<-d2)
}

func generate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, number := range numbers {
			channel <- number
		}
		close(channel)
	}()

	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()

	return channel
}
