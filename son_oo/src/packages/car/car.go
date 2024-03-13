package car

type Car struct {
	Name  string
	Color string
}

func (c Car) Start() string {
	return "The Car " + c.Name + " Color: " + c.Color + " started"
}
