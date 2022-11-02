package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"CustomizedName"`
	Year  int    `json:"-"` // This field will not be marshalled
	Color string `json:"CustomizedColor"`
}

func main() {
	car := Car{"BMW", 2010, "Black"}
	result, _ := json.Marshal(car)

	fmt.Println(string(result))
}
