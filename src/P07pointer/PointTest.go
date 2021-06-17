package P07pointer

import "fmt"

func PointTest01() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

//测试引用传递
func PointTest02() {
	var a int = 100
	var b int = 200

	fmt.Printf("交换前: a 的值： %d\n\n", a)
	fmt.Printf("交换前: b 的值： %d\n\n", b)
	swap(&a, &b)

	fmt.Printf("交换后: a 的值： %d\n\n", a)
	fmt.Printf("交换后: b 的值： %d\n\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
