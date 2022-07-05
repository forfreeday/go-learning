package p26httpweb

import (
	"fmt"
	"net/http"
	"strings"
)

func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                   // 解析参数
	fmt.Println(r.Form)             // 在服务端打印请求参数
	fmt.Println("URL:", r.URL.Path) // 请求 URL
	fmt.Println("Scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "你好，嗯哼！") // 发送响应到客户端
}
