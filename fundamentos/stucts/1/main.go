package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Corola", 2017, "White"}
	//car2 := Car{"BMW 320", 2017, "BLACK"}

	//fmt.Println(car1.Name, car1.Year, car1.Color)
	//fmt.Println(car2.Name, car2.Year, car2.Color)
	fmt.Println(car1.info())
}
