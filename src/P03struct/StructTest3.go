package P03struct

import "fmt"

type Student2 struct {
	Name string
	Age  int
}

func (student Student2) TestMethod() {
	student.Name = "ttttttttttt"
	fmt.Println("测试结构体方法绑定: ", student)
}

func Test4() {
	stud := Student2{
		Name: "啊哈",
		Age:  18,
	}

	stud2 := Student2{
		Name: "哦吼",
		Age:  20,
	}

	stud.TestMethod()
	stud2.TestMethod()
}

func (student *Student2) TestMethod2() {
	var i = &student.Age
	fmt.Println(&student.Age)
	fmt.Println(*i)

	(*student).Name = "cccc"
	fmt.Println("测试结构体方法绑定: ", student)
}

func Test5() {
	stud := Student2{
		Name: "啊哈",
		Age:  18,
	}

	stud2 := Student2{
		Name: "哦吼",
		Age:  20,
	}

	stud.TestMethod2()
	stud2.TestMethod2()
}

func Test6() {
	stud := &Student2{
		Name: "啊哈",
		Age:  18,
	}

	stud2 := &Student2{
		Name: "哦吼",
		Age:  20,
	}

	stud.TestMethod()
	stud2.TestMethod2()

	fmt.Println(stud)
	fmt.Println(stud2)
}
