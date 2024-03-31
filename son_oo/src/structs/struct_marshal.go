package main

import (
	"encoding/json"
	"fmt"
)

// type Car struct {
// 	Name  string
// 	Year  int
// 	color string
// }

type SuperCar struct {
	Car
	CanFly bool
}

// func (c Car) info() string {
// 	return fmt.Sprintf("Car: %s Year: %d Color: %s", c.Name, c.Year, c.Color)
// }

func main() {
	car2 := Car{"Fusca", 2030, "Blue"}
	result, _ := json.Marshal(car2)
	fmt.Println(string(result))
}
