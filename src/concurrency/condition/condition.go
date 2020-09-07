package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Inside Main")
	c := sync.NewCond(&sync.Mutex{}) // Initialize with initial mutex which will be used
	c.L.Lock()
	for conditionTrue() == true {
		c.Wait() //the call to Wait automatically calls Unlock on the Locker when entered.
		//   Here we wait to be notified that the condition has occurred.
		//  This is a blocking call and the goroutine will be suspended.
		//  A few other things happen when you call Wait: upon entering Wait, Unlock is called on the Cond variable’s Locker,
		//  and upon exiting Wait, Lock is called on the Cond variable’s Locker.
	}
	c.L.Unlock()
	fmt.Println("done part One")
	// Producer Consumer Example using Go

	condition := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10) //slice with a length of zero. & we instantiate it with a capacity of 10.

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		condition.L.Lock() // You can have the next statement as defer condition.L.Unlock but i think
		// you would want to signal also. Also need to check what happens if we signal without releaseing lock
		// Most probably it will throw error saying all the go routines are sleeping.
		queue = queue[1:]
		fmt.Println("Removed element from queue")
		condition.L.Unlock()
		condition.Signal() // Consumer calls this
	}

	for i := 0; i < 10; i++ {
		condition.L.Lock()
		for len(queue) == 2 {
			condition.Wait() // Producer is waiting for stuff to be consumed. Ideally this should be equal to the length of the queue
			// So if the consumer is slow the producer waits
			// TODO implement one with logic above and with go routine out of the loop
		}

		fmt.Println("Adding to Queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // This can even be moved out of the loop
		// Need to better understand the consequences of moving it out
		condition.L.Unlock()
	}
}

func conditionTrue() bool {
	time.Sleep(1 * time.Second)
	return false
}
