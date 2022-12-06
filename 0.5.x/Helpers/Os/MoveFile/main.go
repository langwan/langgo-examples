package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
)

func main() {
	wn, err := helper_os.MoveFile("./ls.bk", "./ls.bk2")

	if err != nil {
		panic(err)
	}
	fmt.Printf("move file %d bytes\n", wn)
}
