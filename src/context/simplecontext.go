package main

import (
	"context"
	"fmt"
)

// MyInt is stupid
type MyInt int

func main() {
	fmt.Println("GO")
	var in MyInt = 1
	context.WithValue(context.Background(), in, 2)

}
