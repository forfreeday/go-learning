package main

import (
	"chart02"
	"fmt"
	"pstruct"
	_ "pstruct"
)

func main() {
	Chart02Test()
}

func Chart02Test() {
	tree := chart02.TreeImpl{10}
	fmt.Println("tree", tree.Call())

	structTest := pstruct.Vertex{X: 10.0, Y: 20.0}
	fmt.Println("structTest", structTest.Abs())
}
