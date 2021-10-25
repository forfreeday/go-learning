package main

import "fmt"

//for
func test4() {
	var sliceTest = make([]int, 5, 20)
	sliceTest[0] = 111
	sliceTest[1] = 222
	sliceTest[2] = 333
	sliceTest[3] = 444
	sliceTest[4] = 555
	for i := 0; i < len(sliceTest); i++ {
		fmt.Printf("for i sliceTest[%d]=%d\n", i, sliceTest[i])
	}

	for i, v := range sliceTest {
		fmt.Printf("for range sliceTest[%d]=%d\n", i, v)
	}
}

func test3() {
	var sliceTest1 = []string{"嗯哼", "啊哈", "哦吼"}
	var sliceTest2 []string = []string{"嗯哼", "啊哈", "哦吼"}
	sliceTest5 := []string{"嗯哼", "啊哈", "哦吼"}
	fmt.Println("sliceTest 数据", sliceTest1)
	fmt.Println("sliceTest 数据", sliceTest2)
	fmt.Println("sliceTest 数据", sliceTest5)
}

func test2() {
	var sliceTest = make([]int, 10, 20)
	sliceTest[0] = 111
	sliceTest[1] = 222
	fmt.Println("sliceTest 数据", sliceTest)
}

func test1() {
	var arr [5]int = [...]int{1, 2, 3, 4, 5}
	//初始化切片，最后下标为3但不包含3
	sliceTest := arr[1:3]

	fmt.Println("intArr=", arr)
	fmt.Println("sliceTest 数据", sliceTest)
	fmt.Println("sliceTest 个数=", len(sliceTest))
	//动态可扩展，一般是长度的两倍，java的容器策略
	fmt.Println("sliceTest 容量=", cap(sliceTest))

	//验证 sliceTest 修改的是原数组的数据
	sliceTest[0] = 99999
	fmt.Printf("arr=%v\n", arr)

}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}
