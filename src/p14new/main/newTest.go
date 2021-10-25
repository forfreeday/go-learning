package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("num1 类型%T, num1的值%v, num1的地址%v\n", num1, num1, &num1)

	num2 := new(int) //返回一个 *int 类型指针
	fmt.Printf("num2 类型%T, num2的值%v, num2的地址%v", num2, num2, &num2)
}
