package p26httpweb

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
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

// RunShell 执行 shell 操作
func RunShell(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Fprintf(w, "combined out:\n%s\n", string(output)) // 发送响应到客户端
}

type Response struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}

func RespJson(w http.ResponseWriter, r *http.Request) {
	response := &Response{
		"in_channel",
		"调用daily build 成功",
	}
	result, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))
	fmt.Fprintf(w, string(result)) // 发送响应到客户端
}
