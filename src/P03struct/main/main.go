package main

import (
	. "P02interface"
	"P03struct"
	"fmt"
)

func main() {
	//test1()
	test2()
}

func test1() {
	tree := TreeImpl{PhoneNumber: 10}
	fmt.Println("tree", tree.Call())
}

func test2() {
	//P03struct.StructTest1()
	//P03struct.StructTest2()
	//P03struct.StructTest3()
	P03struct.StructTest31()

}
