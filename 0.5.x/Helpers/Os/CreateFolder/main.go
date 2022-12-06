package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
)

func main() {
	err := helper_os.CreateFolder("./newfolder", true)
	if err != nil {
		panic(err)
	}
	err = helper_os.CreateFolder("./newfolder", true)
	if err != nil {
		panic(err)
	}
	fmt.Println("create newfolder ok")

	err = helper_os.CreateFolder("./newfolder2", false)
	if err != nil {
		fmt.Println("create newfolder2 failed")
		return
	}
	err = helper_os.CreateFolder("./newfolder2", false)
	if err != nil {
		fmt.Println("create newfolder2 failed")
		return
	}
	fmt.Println("create newfolder2 ok")
}
