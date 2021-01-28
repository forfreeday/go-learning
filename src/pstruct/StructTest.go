package pstruct

import "fmt"

type Vertex struct {
	X float32
	Y float32
}

func (v *Vertex) Abs() float32 {
	return v.X * v.Y
}

func main() {
	test := Vertex{10, 20}
	fmt.Print(test.Abs())
}
