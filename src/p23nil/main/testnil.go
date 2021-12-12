package main

import "log"

func test1() {
	a := (*interface{})(nil)
	log.Printf("%v\n", a == nil)
	var c interface{}

	c = (*interface{})(nil)
	log.Printf("%v\n", c == nil)
	log.Printf("%v\n", c)
}

// Text ------------------------------
type Text interface {
	Abs()
}

type noText struct{}

func (noText) Abs() {}

var _ Text = (*noText)(nil)

func text2() {
	log.Printf("%v\n", Text.Abs == nil)
}

func main() {
	//test1()

	text2()
}
