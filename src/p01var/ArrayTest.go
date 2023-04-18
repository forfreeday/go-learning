package p01var

import "fmt"

func pointTest(arr *[3]string) {
	//标准写法
	(*arr)[0] = "piu piu piu "
	//go自动优化
	arr[1] = "bui bui bui"
	fmt.Printf("(*arr)[0] 地址=%p\n", &(*arr)[0])
	fmt.Printf("arr[0] 地址=%p\n", &arr[0])
}

func ArrayTest6() {
	hens := [...]string{"嗯哼", "啊哈", "哦吼"}
	pointTest(&hens)
	fmt.Printf("hens[0]=%s\n", hens[0])
	fmt.Printf("hens[1]=%s\n", hens[1])
}

func ArrayTest5() {
	hens := [...]string{"嗯哼", "啊哈", "哦吼"}

	for index, value := range hens {
		//这种就类似于 foreach
		fmt.Printf("index=%d, value=%s\n", index, value)
		//fmt.Printf("hens=%s\n", hens[index])
	}
}

func ArrayTest4() {
	var arr [4]int
	fmt.Println(arr)
	fmt.Printf("arr[0]的地址=%#x\n", &arr[0])
	fmt.Printf("arr[1]的地址=%#x\n", &arr[1])
	fmt.Printf("arr[2]的地址=%#x\n", &arr[2])
	fmt.Printf("arr[2]的地址=%#x\n", &arr[3])
}

func ArrayTest3() {
	var hens [6]int
	hens[0] = 111
	hens[1] = 222
	hens[2] = 333
	//越界
	//hens[9] = 999
	for i := 0; i < len(hens); i++ {
		fmt.Printf("打印数组i=%d, 数据=%d\n", i, hens[i])
	}

}

func ArrayTest() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
func ArrayTest2() {
	var i, j, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	/* 输出每个数组元素的值 */
	for j = 0; j < 5; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
}

func ArrayTest7() {
	balance := [10]string{"stress-bot", "deploy", "branch", "test1", "test2"}

	fmt.Printf("length: %d\n", len(balance))
	fmt.Printf("param is: %s", balance[2:])
	for i := 0; i < len(balance); i++ {

	}
}
