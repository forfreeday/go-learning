package main

import (
	"golearning/src/P06import/hello/impl"
	_ "golearning/src/P06import/hello/impl"
)

func main() {
	//imp.Print() 编译报错，说：undefined: imp
	impl.Print()
}
