package main

import (
	"fmt"
	log "golearning/src/p22assetion"
)

func main() {
	//test()
	//test1()
	//test2()
	log.NewNopLogger()
}

func test() {
	var a interface{}

	fmt.Println("Where are you,Jonny?", a.(string))
}

func test1() {
	var a interface{}
	value, ok := a.(string)
	if !ok {
		fmt.Println("It's not ok for type string")
		return
	}
	fmt.Println("The value is ", value)
}

func test2() {
	var t interface{}
	t = functionOfSomeType()
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}

func functionOfSomeType() interface{} {
	return 3333
}

type Point struct {
	x int
	y int
}

func test4() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	// 这里未经过断言，是不可能转换为具体类型的
	b = a.(Point)
	fmt.Println(b)
}

