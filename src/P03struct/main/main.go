package main

import (
	"fmt"
	. "golearning/src/P02interface"
	"golearning/src/P03struct"
	"golearning/src/P03struct/aa/bb"
)

func main() {
	//test1()
	//test2()
	//test3()
	P03struct.Test4()
	P03struct.Test5()
	P03struct.Test6()
}

func test1() {
	tree := TreeImpl{PhoneNumber: 10}
	fmt.Println("tree", tree.Call())
}

func test2() {
	P03struct.StructTest1()
	P03struct.StructTest2()
	P03struct.StructTest3()
	P03struct.StructTest31()
	P03struct.StructTest4()
}

func test3getStudent() *bb.Student {
	student := &bb.Student{
		Name: "test",
		Age:  99,
	}
	return student
}

func test3() {
	student := test3getStudent()
	fmt.Println(student)
}
