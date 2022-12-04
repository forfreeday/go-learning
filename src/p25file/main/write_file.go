package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	os "os"
)

func Write() {
	var buffer bytes.Buffer

	for i := 0; i < 100; i++ {
		buffer.WriteString("this is line " + fmt.Sprintf("%d", i) + "\n")
	}

	if err := ioutil.WriteFile("/Users/liukai/testFile", buffer.Bytes(), 0644); err != nil {
		fmt.Println(err)
	}
}

func Write2() {
	filePath := "/Users/liukai/testWrite.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("open file err %v\n", err)
		return
	}

	str := "test write\n"
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString(str)
	}
	// 真正写入到碰盘，和java一样的缓冲操作
	write.Flush()
}

func ReadThanWrite() {
	readFile, err := ioutil.ReadFile("/Users/liukai/read.txt")
	if err != nil {
		fmt.Println("read file err =%v", err)
	}
	err = ioutil.WriteFile("/Users/liukai/write.txt", readFile, 0666)
	if err != nil {
		fmt.Println("write file err=%v", err)
		return
	}
}

func PathExists() (bool, error) {
	_, err := os.Stat("/Users/liukai/read.txt")
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	//Write2()
	//ReadThanWrite()
	exists, err := PathExists()
	if err != nil {
		return
	}
	fmt.Println("result=%v", exists)
}
