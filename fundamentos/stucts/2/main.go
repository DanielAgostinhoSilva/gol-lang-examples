package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	Name   string
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Corola", 2017, "White"}
	scar := SuperCar{
		Car: Car{"Fusca", 2030, "Blue"},
		//Car: car1,
		CanFly: true,
		Name:   "Fusca Turbo",
	}
	fmt.Println(car1.info())
	fmt.Println(scar.Name)
}
