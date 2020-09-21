package main

import (
	"fmt"
)

type hello struct {
	i int32
}

func main() {

	h1 := &hello{i: 21}
	h2 := &hello{i: 22}

	var h3 *hello = &hello{}
	fmt.Printf(" h1 = %v  and h2 = %v \n ", &h1, &h2)
	fmt.Printf(" h1 = %v  and h2 = %v \n ", *h1, *h2)
	fmt.Printf(" h1 = %v  and h2 = %v \n ", h1, h2)
	fmt.Printf(" h1 = %v  and h2 = %v \n ", *&h1, *&h2)
	fmt.Printf("h3 = %v", &h3)
}
