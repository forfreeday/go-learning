package p08string

import "fmt"

func TestString() {

	var c1 byte = 'a'
	var c2 byte = '0' //字符0

	//直接输出是 ASCII 码
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)

	//格式化输出
	fmt.Printf("format=%c\n", c1)

	//汉字是 unicode，超过 byte 256
	var c3 int = '嗯' //21999
	fmt.Printf("汉字format=%c, unicode=%d\n", c3, c3)

	var c4 int = 21999
	fmt.Printf("unicode转汉字=%c unicode=%d\n", c4, c4)

	var address string = "Hello World 北京"
	fmt.Println(address)
}
