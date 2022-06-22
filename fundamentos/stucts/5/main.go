package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	Color string
}

func main() {

	var car Car

	data := []byte(`{ 
		"Name" : "Gol",
		"Year" : 2017,
		"Color" : "Yellow"
	}`)

	json.Unmarshal(data, &car)

	//fmt.Println(result)
	fmt.Println(car)
}
