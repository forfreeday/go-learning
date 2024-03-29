package main

import "fmt"

func test8() {
	testString := "test@aaa.com"
	arr := []rune(testString)
	arr[0] = '北'
	testString = string(arr)
	fmt.Println(testString)

}

//string
func test7() {
	testString := "test@aaa.com"
	arr := []byte(testString)
	arr[0] = 'z'
	testString = string(arr)
	fmt.Println(testString)

}

//copy
func test6() {
	var a = []int{11, 22, 33, 44}
	var sliceTest = make([]int, 1)
	copy(sliceTest, a)
	fmt.Println(sliceTest)
}

//append
func test5() {
	sliceTest := []int{11, 22, 33, 55}
	sliceTest = append(sliceTest, 44, 66, 77)
	fmt.Println(sliceTest)
}

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
	sliceTest1 := arr[0:3]
	sliceTest2 := arr[1:3]

	fmt.Println("intArr=", arr)
	fmt.Println("sliceTest1 数据", sliceTest1)
	fmt.Println("sliceTest2 数据", sliceTest2)

	fmt.Println("sliceTest1 个数=", len(sliceTest1))
	//动态可扩展，一般是长度的两倍，java的容器策略
	fmt.Println("sliceTest1 容量=", cap(sliceTest1))

	//验证 sliceTest 修改的是原数组的数据
	sliceTest1[0] = 99999
	fmt.Printf("arr=%v\n", arr)

}

func main() {
	test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
}
