package main

import (
	. "P02interface"
	"fmt"
)

func main() {
	Chart02Test()
}

func Chart02Test() {
	tree := TreeImpl{PhoneNumber: 10}
	fmt.Println("tree", tree.Call())
}
