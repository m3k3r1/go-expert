package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\n Year: %d\n Color: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"BMW", 2010, "Black"}
	car2 := Car{"Mercedes", 2015, "White"}

	fmt.Println(car1)
	fmt.Println(car1.Name)
	fmt.Println(car2)
	fmt.Println(car2.Year)

	fmt.Println(car1.info())
	fmt.Println(car2.info())
}
