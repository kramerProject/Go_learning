package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s Year: %d Color: %s", c.Name, c.Year, c.Color)
}

func main() {
	car := Car{}
	car.Name = "FORD"
	car.Year = 2022
	car.Color = "Black"
	fmt.Println(car.info())
}
