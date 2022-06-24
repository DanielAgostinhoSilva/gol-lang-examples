package main

import "fmt"

// Pipeline pattern
func main() {
	numbers := generate(2, 4, 6)
	result := devide(numbers)

	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
}

func generate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, number := range numbers {
			channel <- number
		}
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
