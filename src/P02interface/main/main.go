package main

import (
	"fmt"
	"golearning/src/P02interface"
)

func test1() {
	tree := P02interface.TreeImpl{10}
	fmt.Println("tree", tree.Call())
}

func test2() {
	computer := P02interface.Computer{}
	phone := P02interface.Phone{}
	camera := P02interface.Camera{}
	computer.Working(phone)
	computer.Working(camera)
}

func main() {
	//test1()
	test2()
}
