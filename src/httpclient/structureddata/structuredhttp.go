package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

func (svr server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "{\"message\": \"hello world\"}"

	w.Write([]byte(msg))
}

func main() {
	fmt.Println("Hello Server")
	log.Fatal(http.ListenAndServe(":8080", server{}))

}
