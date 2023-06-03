package util

import (
	"fmt"
	"os"
)

// copyFile 文件拷贝
func copyFile(source, destination string) {
	input, err := os.ReadFile(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile(destination, input, 0644)
	if err != nil {
		fmt.Println(destination)
		fmt.Println(err)
		return
	}
}
