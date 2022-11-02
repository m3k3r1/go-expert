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
	var car Car

	data := []byte(`{"Name":"BMW","Year":2010,"Color":"Black"}`)
	_ = json.Unmarshal(data, &car)
	fmt.Println(car)

}
