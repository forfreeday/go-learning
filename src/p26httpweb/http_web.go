package p26httpweb

import (
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

func RunShell(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("ls", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Fprintf(w, "combined out:\n%s\n", string(output)) // 发送响应到客户端
}

func RunDailyBuild(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sh /data/workspace/docker_workspace/new-stest_daily_task.sh ", "solc-0.8.11", ">/data/workspace/docker_workspace/stest_shell.log 2>&1 < /dev/null &")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Fprintf(w, "combined out:\n%s\n", string(output)) // 发送响应到客户端
}
