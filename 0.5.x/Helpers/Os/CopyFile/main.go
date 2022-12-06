package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
)

func main() {
	wn, err := helper_os.CopyFile("./ls.bk", "/bin/ls")
	if err != nil {
		panic(err)
	}
	fmt.Printf("copy file %d\n bytes", wn)
}
