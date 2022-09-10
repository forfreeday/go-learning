package main

import (
	"golearning/src/p26httpweb"
	"log"
	"net/http"
)

func main() {

	//http.HandleFunc("/shell", RunDailyBuild2)
	//err := http.ListenAndServe(":9091", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe DailyBuild: ", err)
	//}

	http.HandleFunc("/json", p26httpweb.RespJson)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe Shell: ", err)
	}

	http.HandleFunc("/shell", p26httpweb.RunShell)
	err = http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe Shell: ", err)
	}

	http.HandleFunc("/", p26httpweb.SayHelloWorld)
	err = http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

//func RunDailyBuild2(w http.ResponseWriter, r *http.Request) {
//
//	cmdStr := "/data/workspace/docker_workspace/new-stest_daily_task.sh solc-0.8.11>/data/workspace/docker_workspace/stest_shell.log 2>&1 "
//	cmd := exec.Command("sh", "-c", cmdStr)
//	_, err := cmd.CombinedOutput()
//	if err != nil {
//		log.Fatalf("cmd.Run() failed with %s\n", err)
//	}
//	fmt.Fprintf(w, "daily build running.") // 发送响应到客户端
//}
