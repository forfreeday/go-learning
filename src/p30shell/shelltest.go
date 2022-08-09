package main

import (
	"fmt"
	"log"
	"os/exec"
)

// 测试shell执行
func main() {
	cmd := exec.Command("ls", "-l")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(output))
}
