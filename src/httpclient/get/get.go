package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r, err := http.Get("http://www.tired.com")
	defer r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
