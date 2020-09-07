package main

import (
	"fmt"
)

type Outer struct {
	name string
	Inner
}

type Inner struct {
	age uint
}

//type InnerInterface interface {
//	GetAge() uint
//}

func (inner Inner) GetAge() uint {
	return inner.age
}

func main() {
	fmt.Println(" Hello behaviour")
	inner := Inner{age: 23}
	outer := Outer{name: "Beaut", Inner: inner}
	fmt.Println(outer)
	fmt.Println(outer.age)
	outer.GetAge()

}
