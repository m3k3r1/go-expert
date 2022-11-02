package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s", c.Name, c.Year, c.Color)
}

func main() {
	car := Car{"Mercedes", 2015, "White"}
	superCar := SuperCar{Car: Car{"Fusca", 2030, "Black"}, CanFly: true}

	fmt.Println(car)
	fmt.Println(car.info())

	fmt.Println(superCar)
	fmt.Println(superCar.info())
}
