package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello in main")
	ch := make(chan string)
	const poolsize = 2
	for i := 0; i < poolsize; i++ {
		go func(p int) {
			for str := range ch {
				time.Sleep(100000)
				fmt.Printf("p %d received signal %s ", p, str)
			}
			fmt.Printf("p %d received signal shutdown\n", p)
		}(i)
	}
	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
		fmt.Println("\nSignal Sent ", w)
	}
	close(ch)
	fmt.Println("Shutdown Signal Sent")
	time.Sleep(2 * time.Second)
	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
}
