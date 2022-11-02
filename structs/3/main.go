package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	Color string
}

func main() {
	car := Car{"BMW", 2010, "Black"}
	result, _ := json.Marshal(car)

	fmt.Println(result)
	fmt.Println(string(result))
}
