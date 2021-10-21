package P10meth

import "fmt"

type Cycle struct {
	Redis int
}

func (cycle *Cycle) CycleTest(p int) int {
	fmt.Printf("cycle=%p\n", cycle)
	fmt.Printf("&cycle=%p\n", &cycle)
	fmt.Printf("*cycle=%d\n", *cycle)
	return cycle.Redis * p
}

func main() {
	var cycle Cycle
	cycle.Redis = 111
	fmt.Printf("main &cycle=%p\n", &cycle)
	sum1 := cycle.CycleTest(20)
	fmt.Println("sum1=", sum1)
}
