package main

import "fmt"

type myInt int

// 为类型 myInt 方法绑定 IsZero
func (mi myInt) IsZero() bool {
	fmt.Println("a=", mi)
	return mi == 0
}

func test3() {
	var a myInt
	a = 10
	fmt.Println(a.IsZero()) // false
}
func main() {
	test3()
}

func test1() {
	var myTest myInt
	myTest = 123
	fmt.Println(myTest)
}

func test2() {
	var a int
	var myTest myInt

	myTest = 123

	a = 2
	b := myInt(a)
	fmt.Println(b) // 2
	fmt.Println(myTest)
}
