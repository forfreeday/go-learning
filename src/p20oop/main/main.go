package main

import (
	"fmt"
)

// 父结构体
type Father struct {
	Name string
	age  int
}

type Mother struct {
	Name string
	age  int
}

// 子结构体
type Son struct {
	Father
	grade string // 年级
}

type LitterSon struct {
	*Father
	*Mother
}

// 为父亲绑定一个方法，父亲的方法儿子也可以继承使用
func (fa Father) eat() {
	fmt.Println(fa.Name, "eat")
}

func (fa *Father) sleep() {
	fmt.Println(fa.Name, "sleep")
}

func test1() {
	son := &Son{}
	// 1、通过Father给父结构体的属性赋值
	son.Father.age = 40
	son.Father.Name = "父亲"
	son.grade = "6年级"
	fmt.Println(son) // &{ {父亲 40} 6年级}

	son2 := &Son{}
	// 2、直接给Father的属性赋值
	(*son2).age = 40
	son2.Name = "son2儿子"
	son2.grade = "3年级"
	fmt.Println(son2) //&{ {son2父亲 40} 3年级}

	son2.eat()
	son2.sleep()
}

func test2() {
	son := &LitterSon{
		&Father{
			Name: "嗯哼",
			age:  11,
		},
		&Mother{
			Name: "啊哈",
			age:  11,
		},
	}
	// 指针的是两个地址，是正确的
	fmt.Println((*son))
	fmt.Println(son.Father, (*son).Mother)
}

func test3() {
	son := Son{
		Father: Father{
			Name: "嗯哼",
			age:  11,
		},
	}

	son2 := Son{Father: Father{Name: "嗯哼", age: 11}}
	// 指针的是两个地址，是正确的
	fmt.Println(son)
}

func main() {
	//test1()
	test2()
}
