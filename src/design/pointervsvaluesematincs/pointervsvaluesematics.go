package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) scaleByVal(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v1 := Vertex{X: 10, Y: 10}
	fmt.Println(v1)
	v1.scale(10)
	fmt.Println(v1)
	v1.scaleByVal(10)
	fmt.Println(v1)
}
