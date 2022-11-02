package main

import "fmt"

func main() {
	var x [10]int

	fmt.Println(x)
	fmt.Println(len(x))

	x[0] = 10
	x[1] = 20
	x[2] = 30

	fmt.Println(x)
	fmt.Println(x[0])

	var y [10]string

	fmt.Println(y)

	z := [5]int{1, 2, 3, 4, 5}
	fmt.Println(z)
}
