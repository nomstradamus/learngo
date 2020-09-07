package main

import (
	"fmt"
)

type Size int

const (
	small Size = iota
	medium
	large
)

type MyStruct struct {
	name string
	size Size
	s    int64
}

func main() {
	mystruct := MyStruct{
		name: "hello",
		size: small,
		s:    32,
	}
	fmt.Printf("%d\n", small)
	fmt.Println("MyStruct", mystruct)
	myarr := []string{"one", "two", "three"}
	fmt.Println(" Array ", myarr)
	// format fo loop in go 1
	for i, s := range myarr {
		fmt.Println(i, s)
		s = "four"
	}
	fmt.Println(" Array ", myarr)
	// format for loop in go 2
	for i := 0; i < len(myarr); i++ {
		myarr[i] = "three"
	}
	fmt.Println(" Array ", myarr)
	x := 0
	// Infinite loop with break statement
	for true {
		x = x + 1
		fmt.Println(x)
		if x == 3 {
			break
		}
	}

}
