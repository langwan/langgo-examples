package main

import helper_os "github.com/langwan/langgo/helpers/os"

func main() {
	helper_os.TouchFile("./langgo", true, false)
	helper_os.TouchFile("./foo/foo", true, true)
}
