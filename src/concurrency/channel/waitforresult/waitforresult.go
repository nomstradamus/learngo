package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func main() {
	// I am creating a channel which will contain value of type int
	ch := make(chan int)
	// I am creating a function which I will call later
	// as a go routine ( basically it will run in different thread while main continues to execute)
	fun := func() {
		time.Sleep(1000)
		color.Red("Inside")
		ch <- 12

	}
	go fun()
	color.Green(" This will be executed before go routine ")
	val := <-ch
	color.Green(" This will be executed after go routine ")
	fmt.Printf("%d\n", val)
	color.Cyan("Prints text in cyan.")

}
