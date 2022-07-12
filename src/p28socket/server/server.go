package main

import (
	"fmt"
	"net" // 做 socket编程需要
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)

		//fmt.Println("server waiting for client: " + conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("server read err: ", err)
			return
		}
		fmt.Println(string(buf[:n])) // 注意这里要切片 [:n], 不然会输出 1024个字节
	}
}

func main() {
	fmt.Println("开始监听...")
	// 指定网络协议以及本地监听端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen error")
		return
	}

	defer listen.Close()

	for {
		fmt.Println("waiting for client...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err = ", err)
		} else {
			fmt.Println("Accept suc con = ", conn)
		}
		go process(conn)
	}
}
