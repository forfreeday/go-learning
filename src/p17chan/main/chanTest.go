package main

import (
	"fmt"
	"time"
)

func go11(message chan string) {
	message <- "go1 hello1"
	message <- "go1 hello2"
	message <- "go1 hello3"
	message <- "go1 hello4"
}
func go22(message chan string) {
	time.Sleep(1 * time.Second)
	str := <-message
	str = "go2 get chan message: " + str
	message <- str
	//关闭chan
	close(message)
}

func test7() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		// 从3循环到0
		for i := 3; i >= 0; i-- {
			// 发送3到0之间的数值
			ch <- i
			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()
	// 遍历接收通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)
		// 当遇到数据0时, 退出接收循环
		if data == 0 {
			break
		}
	}
}

func test4() {
	var message = make(chan string, 3)
	go go11(message)
	go go22(message)
	time.Sleep(2 * time.Second)
	//range读取
	for val := range message {
		fmt.Println(val)
	}
}

func test6() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	// 等待匿名goroutine
	<-ch
	fmt.Println("all done")
}

//验证阻塞
func test5() {
	var testChan = make(chan string)
	//发送的阻塞的，所在需要使用协程
	go func() {
		//testChan <- "send test message"
	}()
	data, ok := <-testChan
	if ok {
		fmt.Println(data)
	}

}

func test3() {
	var testChan = make(chan string)
	//发送的阻塞的，所在需要使用协程
	go func() {
		//testChan <- "send test message"
	}()
	str := <-testChan

	fmt.Println(str)
}

//创建 channel
func test1() {
	////使用前必须先 make
	//var testChan = make(chan string)
	////创建一个 int类型，容量100的chan
	//var testChanInt = make(chan int, 100)
	//
	//
	//ch1 := make(chan int)                 // 创建一个整型类型的通道
	//ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式
	//
	//type Equip struct{ /* 一些字段 */ }
	//ch2 := make(chan *Equip)             // 创建Equip指针类型的通道, 可以存放*Equip
}

func test2() {
	ch1 := make(chan int) // 创建一个整型类型的通道
	ch1 <- 100
	fmt.Println(<-ch1)
}

func main() {
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
}
