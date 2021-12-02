package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-a")
	//标准输出
	cmd.Stdout = os.Stdout
	//错误输出
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to call cmd.Run(): %v", err)
	}
}
