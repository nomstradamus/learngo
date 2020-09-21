package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("In main")
	ch := make(chan struct{})

	fun := func() {
		time.Sleep(1 * time.Second)
		close(ch)
		fmt.Println("Closed channel")
	}

	listeners := func(v int) {
		fmt.Println("Waiting for channel to close ", v)
		<-ch
		fmt.Println("Channel closed event : ", v)
	}

	go listeners(1)
	go listeners(2)

	go fun()
	_, wd := <-ch
	fmt.Println(" Received Signal In main : ", wd)
	time.Sleep(time.Second)

}
