package main

import "fmt"

type Names []interface {
}

func (n *Names) load() {
	*n = Names{
		"Daniel",
		"Sandra",
	}
}

func (n Names) printName() {
	fmt.Println(n)
}

func main() {
	var names Names
	names.load()
	names.printName()
}
