package main

import (
	//. "P02interface"
	//"P03struct"
	"P03struct/aa/bb"
	"fmt"
)

func main() {
	//test1()
	//test2()
	student := test3()
	fmt.Println(student)
}

//
//func test1() {
//	tree := TreeImpl{PhoneNumber: 10}
//	fmt.Println("tree", tree.Call())
//}
//
//func test2() {
//	//P03struct.StructTest1()
//	//P03struct.StructTest2()
//	//P03struct.StructTest3()
//	//P03struct.StructTest31()
//	P03struct.StructTest4()
//}

func test3() *bb.Student {
	student := &bb.Student{
		Name: "test",
		Age:  99,
	}
	return student
}
