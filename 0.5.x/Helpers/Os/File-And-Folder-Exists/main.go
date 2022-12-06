package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
)

func main() {
	ok := helper_os.FileExists("/bin/ls")
	fmt.Println(ok)

	ok = helper_os.FileExists("/bin")
	fmt.Println(ok)

	ok = helper_os.FolderExists("/bin/ls")
	fmt.Println(ok)

	ok = helper_os.FolderExists("/bin")
	fmt.Println(ok)
}
