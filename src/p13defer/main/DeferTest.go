package main

import "fmt"

func sum(n1 int, n2 int) int {
	defer fmt.Printf("n1=%d, n2=%d\n", n1, n2)
	res := n1 + n2
	defer fmt.Printf("res1=%d\n", res)
	return res
}

func sum2(n1 int, n2 int) int {
	defer fmt.Println("n1=", n1)
	defer fmt.Println("n2=", n2)
	n1++
	n2++
	res := n1 + n2
	defer fmt.Printf("res1=%d\n", res)
	return res
}

func main() {
	//res := sum(10, 20)
	//fmt.Printf("res2=%d", res)

	res2 := sum2(10, 20)
	fmt.Printf("res2=%d", res2)
}
