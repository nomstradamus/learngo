package main

import (
	"fmt"
)

type substructinterface interface {
	hola()
}

type innerstruct struct {
	a int32
}

func (inner *innerstruct) hola() {
	fmt.Println("Inner struct hola")
}

type outerstruct struct {
	innerstruct
}

func main() {

	fmt.Println("hello there in main")
	outer := outerstruct{}
	outer.hola()
	fmt.Println(outer.a)
}
