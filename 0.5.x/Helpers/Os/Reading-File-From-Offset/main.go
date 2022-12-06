package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
)

func main() {
	data, err := helper_os.ReadFileAt("/usr/share/man/man1/open.1", 50, 20)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
