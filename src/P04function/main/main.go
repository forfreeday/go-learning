package main

import (
	"fmt"
	"golearning/src/P04function"
)

func test1() {
	result1, result2 := P04function.FunctionTest2(1, 2)
	fmt.Println(result1, result2)
}

//func main() {
//	//test1()
//	callback(1, Add)
//}

func callback(y int, f func(int, int)) {
	// 和lambda一样，只要在真正需要的时候调用一下
	f(y, 2) // this becomes Add(1, 2)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
