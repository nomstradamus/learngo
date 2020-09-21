package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello from go")
	fmt.Printf("%d", runtime.NumCPU())
}
