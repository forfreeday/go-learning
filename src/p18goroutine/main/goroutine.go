package main

import "fmt"

func loop() {
	fmt.Println("-------start---------")
	for i := 0; i < 50; i++ {
		fmt.Printf("loop: %d\n", i)
	}
}

func main() {
	go loop() // 启动一个goroutine
	for i := 0; i < 50; i++ {
		fmt.Printf("main: %d \n", i)
	}
}
