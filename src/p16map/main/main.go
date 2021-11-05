package main

import (
	"fmt"
	"sort"
)

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
	client["test3"] = "哦吼"

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

func test8() {
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	//排序
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}

}

// 切片测试
func test7() {
	// 块明一个切片
	var client []map[string]string
	client = make([]map[string]string, 2)
	if client[0] == nil {
		//这里没有 []
		client[0] = make(map[string]string, 2)
		client[0]["test1-a"] = "嗯哼"
		client[0]["test1-b"] = "嗯哼"
	}

	if client[1] == nil {
		client[1] = make(map[string]string, 2)
		client[1]["test2"] = "啊哈"
	}

	// 这里放开可以测试越界
	//if client[2] == nil {
	//	client[2] = make(map[string]string, 2)
	//	client[2]["test2"] = "啊哈"
	//}

	newClient := map[string]string{
		"test3": "哦吼",
	}

	client = append(client, newClient)
	fmt.Println(client)
}

type student struct {
	Name string
	age  int
}

func test9() {
	var stu map[string]student
	stu = make(map[string]student, 2)

	student1 := student{
		Name: "嗯哼",
		age:  99,
	}
	student2 := student{
		Name: "啊哈",
		age:  11,
	}

	stu["test1"] = student1
	stu["test2"] = student2
	fmt.Println(stu)
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	test9()
}
