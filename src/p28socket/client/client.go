package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 聊天客户端
func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		//fmt.Println("con: ", conn)
	}
	for {
		// 从终端读取一行数据
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err: ", err)
		}
		line = strings.Trim(line, " \r\n") // 小心返回之后要赋值
		if line == "exit" {
			fmt.Println("break")
			return
		}
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err: ", err)
		}
		//fmt.Printf("客户端发送了 %d 字节的数据\n", n)
	}
}
