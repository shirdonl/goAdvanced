package main

import (
	"fmt"
	"os"
)

func main() {
	err := returnsError()
	if err != nil {
		fmt.Printf("操作失败: %v", err)
	}
}

func returnsError() error {
	var err *os.PathError = nil
	// …
	return err
}

//操作失败: <nil>
