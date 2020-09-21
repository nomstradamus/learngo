package main

import (
	"fmt"
)

type Hello interface {
	hello()
}

type HelloInEnglish struct {
}

func main() {

	fmt.Println("Hello In main ")
	eng := HelloInEnglish{}
	fmt.Println(" eng ", eng)
	eng.hello()

}

func (h *HelloInEnglish) hello() {
	fmt.Println("Hello from English ")
}
