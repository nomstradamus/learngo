package main

import (
	"fmt"
)

type flier interface {
	fly()
}

type Airplane struct {
	sound string
}

type SpaceShip struct {
	instruction string
}

func (ap *Airplane) fly() {
	fmt.Println("Airplane fly away")
	fmt.Println(" -> ", ap.sound)
}

func (ss *SpaceShip) fly() {
	fmt.Println("Spaceship fly away")
	fmt.Println(" ->", ss.instruction)
}

func main() {
	ss := SpaceShip{
		instruction: "I go far",
	}
	ap := Airplane{
		sound: "zzzz",
	}
	ss.fly()
	ap.fly()

}

func start(f flier) {
	f.fly()
}
