package main

import (
	"time"

	"github.com/fatih/color"
)

func main() {
	color.Red("Inside Main")
	ch := make(chan int)

	fun := func() {
		time.Sleep(1 * time.Second)
		color.Green("GO CO ROUTINE")
		ch <- 10

	}
	go fun()
	color.Red("Inside Main Again")
	<-ch
	color.Red("Got result back from go coroutine")
}
