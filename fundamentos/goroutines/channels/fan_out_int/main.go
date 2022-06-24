package main

import "fmt"

func main() {
	c := generate(6, 10)

	d1 := devide(c)
	d2 := devide(c)

	fmt.Println(<-d1)
	fmt.Println(<-d2)
}

func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, n := range numbers {
			channel <- n
		}
		close(channel)
	}()
	return channel
}

func devide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()
	return channel
}
