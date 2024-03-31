package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"name"`
	Year  int    `json:"year"`
	Color string `json:"color"`
}

func main() {
	var carro Car
	data := []byte(`{"name": "Gol", "year": 2017: "color": "Blue"}`)
	json.Unmarshal(data, &carro)
	fmt.Println(carro)
}
