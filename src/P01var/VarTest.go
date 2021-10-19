package P01var

import (
	"fmt"
	"unsafe"
)

func VarTest01() {
	var i int
	i = 100
	fmt.Println(i)
}

func VarTest02() {
	var i = 200
	fmt.Println(i)
}

func VarTest03() {
	i := 300
	fmt.Println(i)
}

func VarTest04() {
	var n1 int64 = 100
	fmt.Printf("n1 的数据类型是：%T，占用 %d 字节", n1, unsafe.Sizeof(n1))
}
