package main

import (
	f "fmt"
	"sync"
	"time"
)

func main() {
	f.Println("Inside main")
	var lock sync.Mutex

	f.Println("lock in main", lock)
	var x uint
	x = 0
	go increment1(&x, &lock)
	go increment2(&x, &lock)
	time.Sleep(5 * time.Second)
	f.Println("Value of x in main ", x)

}

func increment1(x *uint, lock *sync.Mutex) {
	time.Sleep(2 * time.Second)
	f.Println("\n Inside increment 1\n")
	defer lock.Unlock()
	lock.Lock()
	*x = *x + 1
	f.Printf("%d \n", *x)
	f.Printf("%v \n", &x)
}

func increment2(x *uint, lock *sync.Mutex) {
	f.Println("\nInside increment 2\n")
	defer lock.Unlock()
	lock.Lock()
	*x = *x + 1
	f.Printf("%d \n", *x)
	f.Printf("%v \n", &x)
}
