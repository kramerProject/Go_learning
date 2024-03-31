package main

import "fmt"

type vehicle interface {
	start() string
}

type Car struct {
	name string
}

type MotorCycle struct {
	name string
}

func (c Car) start() string {
	return "The car " + c.name + " has started"
}

func (m MotorCycle) start() string {
	return "The motorcycle " + m.name + " has started"
}

func startVehicle(v vehicle) string {
	return v.start()
}

func main() {
	c := Car{"Gol"}
	m := MotorCycle{"Moto"}
	fmt.Println(startVehicle(c))
	fmt.Println(startVehicle(m))
	// m := MotorCycle{"Kawasaki"}
	// fmt.Println(c.startCar())
	// fmt.Println(m.startMoto())
}
