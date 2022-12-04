package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func copyFile(src string, dest string) {
	// 1.打开文件
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer srcFile.Close()

	// 2.读取文件内容
	reader := bufio.NewReader(srcFile)
	file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file error=%v\n", err)
		return
	}
	writer := bufio.NewWriter(file)
	defer file.Close()
	io.Copy(writer, reader)
}

func main() {
	copyFile("/Users/liukai/read.txt", "/Users/liukai/copy.txt")
}
