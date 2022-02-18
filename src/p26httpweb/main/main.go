package main

import (
	"golearning/src/p26httpweb"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", p26httpweb.SayHelloWorld)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
