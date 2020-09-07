package main

import (
	f "fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // waitgroup is just a countdownlatch in java
	wg.Add(2)             // Experiment by specifying values such as 1, 2, 3 and see the results
	// Really important that we send a pointer else we are sending a different value to defer
	// which will not work
	go doSomething(&wg, 1)
	go doSomething(&wg, 2)
	wg.Wait() // Waiting for it to complete
}

func doSomething(wg *sync.WaitGroup, val int) {
	defer wg.Done() // kind of pattern to ensure we always do the Done before we exit

	if val == 2 {
		//	f.Printf("Current Time %v\n", time.Now())
		time.Sleep(2 * time.Second)

		//f.Printf("Current Time %v\n", time.Now())
	}
	f.Printf("Done %v\n", val)
}
