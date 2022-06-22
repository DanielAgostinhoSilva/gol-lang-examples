package main

import (
	"fmt"
	"iniciando_com_go_lang/packages/car"
)

func main() {
	gol := car.Car{Name: "Gol"}

	fmt.Println(gol.Start())
}
