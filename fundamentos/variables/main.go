package main

import "fmt"

func main() {
	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'W'
	f := `Uouuu`

	fmt.Printf("%v : %T\n", a, a)
	fmt.Printf("%v : %T\n", b, b)
	fmt.Printf("%v : %T\n", c, c)
	fmt.Printf("%v : %T\n", d, d)
	fmt.Printf("%v : %T\n", e, e)
	fmt.Printf("%v : %T\n", f, f)
}
