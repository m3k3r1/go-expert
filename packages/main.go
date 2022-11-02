package main

import (
	"fmt"
	"packages/car"
)

func main() {
	car1 := car.Car{Name: "Ford", Color: "Yellow"}
	fmt.Println(car1.Start())
}
