package main

import "fmt"

func test1() {
	var client map[string]string
	client = make(map[string]string, 10)
	client["test1"] = "嗯哼"
	client["test2"] = "啊哈"
	client["test2"] = "哦吼"

	fmt.Println(client)
}

func test2() {
	var client = make(map[string]string, 10)
	client["test1"] = "嗯哼"
	client["test2"] = "啊哈"
	client["test2"] = "哦吼"

	fmt.Println(client)
}

func test3() {
	var client = map[string]string{
		"test1": "嗯哼",
	}
	fmt.Println(client)
}

func test4() {
	var client = make(map[string]string, 10)
	client["test1"] = "嗯哼"
	client["test2"] = "啊哈"
	client["test3"] = "哦吼"

	delete(client, "test1")
	fmt.Println(client)

	//重新make，老map会被回收
	client = make(map[string]string)
}

func test5() {
	var client = make(map[string]string, 10)
	client["test1"] = "嗯哼"
	client["test2"] = "啊哈"
	client["test2"] = "哦吼"

	val, ok := client["test1"]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("没有找到key")

	}

}

func test6() {
	var client = make(map[string]string, 10)
	client["test1"] = "嗯哼"
	client["test2"] = "啊哈"
	client["test2"] = "哦吼"

	for i, v := range client {
		fmt.Printf("key=%v, value=%v", i, v)
	}

}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	test6()
}
