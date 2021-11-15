package main

import "fmt"

// 定义一个函数类型，替泛后续函数类型，函数签名相同即可使用
// 类似于java的泛型，只不过是隐式的声明
type op func(a, b int) int

// add 是 op 的类型
func add(a, b int) int {
	return a + b
}

// sub 是 op 的类型
func sub(a, b int) int {
	return a - b
}

func Operator(fu op, a, b int) int {
	return fu(a, b)
}

func main() {
	add := Operator(add, 111, 222)
	sub := Operator(sub, 444, 333)
	fmt.Println("add=", add)
	fmt.Println("sub=", sub)
}
