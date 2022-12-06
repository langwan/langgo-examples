package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
)

func main() {
	files, err := helper_os.ReadDir("../", true)
	if err != nil {
		return
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
