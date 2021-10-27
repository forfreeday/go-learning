package main

import "fmt"

func test1() {
	var mapTest1 map[string]string
	mapTest1 = make(map[string]string, 10)
	mapTest1["aaa"] = "测试map"
	fmt.Println(mapTest1)
}

func main() {
	test1()
}
