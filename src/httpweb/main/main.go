package main

import (
	"httpweb"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpweb.SayHelloWorld)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
