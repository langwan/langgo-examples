package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
)

func main() {
	info, err := helper_os.GetFileInfo("/bin/ls")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", info)
}
