package main

import (
	"errors"
	"fmt"
)

func test() {
	defer func() {
		error := recover()
		if error != nil {
			fmt.Println("error=", error)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readConf(config string) (err error) {
	if config == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件异常")
	}
}

func test2() {
	err := readConf("config1.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test2 没有继续执行")
}

func main() {
	//test()
	test2()
	fmt.Println("main 执行")
}
