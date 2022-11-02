package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Number int `json:"number"`
	Total  int `json:"total"`
}

func main() {
	conta := Conta{Number: 1, Total: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	raw_json := []byte(`{"Number": 1, "Total": 100}`)
	var contaX Conta
	err = json.Unmarshal(raw_json, &contaX)
	if err != nil {
		panic(err)
	}

	println(contaX.Number)
	println(contaX.Total)
}
