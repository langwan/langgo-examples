package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
)

func main() {
	filename := helper_os.FileNameWithoutExt("/foo/langgo.jpg")
	fmt.Println(filename)
}
