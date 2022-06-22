package main

import (
	"fmt"
)

type vehicle interface {
	start() string
}

type motorcycle struct {
	name string
}

type car struct {
	name string
}

func (c car) start() string {
	return "The car " + c.name + " has been started"
}

func (mc motorcycle) start() string {
	return "The motocycle " + mc.name + " has been started"
}

func startVehicle(v vehicle) string {
	return v.start()
}

func main() {

	c := car{"Gol"}
	m := motorcycle{"XPTO"}

	fmt.Println(startVehicle(c))
	fmt.Println(startVehicle(m))
}
