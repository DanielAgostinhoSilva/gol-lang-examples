package scope

import "fmt"

var y int = 20

func main() {
	x := 10
	fmt.Println(x)
	printY()
	fmt.Printf(z)
	PrintZ()
}

func printY() {
	fmt.Println(y)
}
