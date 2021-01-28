package main

import (
	"P02interface"
	"fmt"
)

func main() {
	tree := P02interface.TreeImpl{10}
	fmt.Println("tree", tree.Call())
}
