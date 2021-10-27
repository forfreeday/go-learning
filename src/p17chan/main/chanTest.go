package main

import (
	"fmt"
	"time"
)

func go11(message chan string) {
	message <- "hello1"
	message <- "hello2"
	message <- "hello3"
	message <- "hello4"
}
func go22(message chan string) {
	time.Sleep(2 * time.Second)
	str := <-message
	str = str + "go"
	message <- str
	//关闭chan
	close(message)
}

func main() {
	var message = make(chan string, 3)
	go go11(message)
	go go22(message)
	time.Sleep(3 * time.Second)
	//range读取
	for val := range message {
		fmt.Println(val)
	}
}
