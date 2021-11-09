package P03struct

import "fmt"

type Student2 struct {
	Name string
	Age  int
}

func (student Student2) TestMethod() {
	fmt.Println("测试结构体方法绑定: ", student)
}

func Test4() {
	stud := Student2{
		Name: "啊哈",
		Age:  18,
	}

	stud.TestMethod()
}
