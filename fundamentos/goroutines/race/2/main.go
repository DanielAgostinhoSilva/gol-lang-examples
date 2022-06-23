package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

var result int
var mutex sync.Mutex

func main() {
	//fmt.Println(runtime.NumCPU())
	waitGroup.Add(2)
	go runProcess("P1", 20)
	go runProcess("P2", 20)
	waitGroup.Wait()

	fmt.Println("Toral result:", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {

		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)

		mutex.Lock()
		result++
		fmt.Println(name, "->", "T", "Partial result:", result)
		mutex.Unlock()
	}
	waitGroup.Done()
}
