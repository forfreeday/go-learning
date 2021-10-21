package P10meth

import "fmt"

type Person struct {
	Name string
}

func (p Person) Test() {
	p.Name = "嗯哼~~"
	fmt.Println(p.Name)
}

func (p Person) Test3(i int, j int) int {
	fmt.Println("p.name=", p.Name)
	return i + j
}
