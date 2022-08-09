package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// 打开、打印、关闭文件
func openFile() {
	// file 文件说明
	// 1. file 文件对象
	// 2. file 指针
	// 3. file 文件句柄
	file, err := os.Open("/Users/liukai/workspaces/temp/go/test.log")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//打印一个地址
	fmt.Printf("file=%v", file)

	err = file.Close()
	if err != nil {
		fmt.Println("Close")
	}
}

// 带缓冲区的 reader，不能读大文件
func buffReader() {
	file, err := os.Open("/Users/liukai/workspaces/temp/go/test.log")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)
	}
}

// 一次性打开一个文件
func readFile() {
	filePath := "/Users/liukai/workspaces/temp/go/test.log"
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File constant: \n%s", file)
}

func main() {
	// readFile()
	// openFile()
	buffReader()
}
