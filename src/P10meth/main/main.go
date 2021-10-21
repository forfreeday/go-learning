package main

import (
	"P10meth"
	"fmt"
)

func main() {
	var p2 P10meth.Person
	p2.Name = "test"
	p2.Test()
	fmt.Println("main() p.Name=", p2.Name)

	sum := p2.Test3(10, 20)
	fmt.Println(sum)

	var cycle P10meth.Cycle
	cycle.Redis = 111
	fmt.Printf("main &cycle=%p\n", &cycle)
	sum1 := cycle.CycleTest(20)
	fmt.Println("sum1=", sum1)

}