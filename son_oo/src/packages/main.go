package main

import (
	"fmt"
	"son_oo/packages/car"
)

func main() {
	car := car.Car{"Nivus", "Black"}
	fmt.Println(car.Start())
}
