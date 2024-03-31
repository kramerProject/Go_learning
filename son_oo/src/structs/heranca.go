package main

import "fmt"

// type Car struct {
// 	Name  string
// 	Year  int
// 	Color string
// }

type SuperCar struct {
	Car
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s Year: %d Color: %s", c.Name, c.Year, c.Color)
}

func main() {
	car := Car{}
	car2 := Car{"Fusca", 2030, "Blue"}
	car.Name = "FORD"
	car.Year = 2022
	car.Color = "Black"
	sCar := SuperCar{car2, true}
	fmt.Println(car.info())
	fmt.Println(sCar.info())
}
