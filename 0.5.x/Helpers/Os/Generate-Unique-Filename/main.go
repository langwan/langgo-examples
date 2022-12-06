package main

import (
	"fmt"
	helper_gen "github.com/langwan/langgo/helpers/gen"
	helper_os "github.com/langwan/langgo/helpers/os"
)

func main() {
	filename, err := helper_os.GenUniqueFilename("/bin/ls", 10, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(filename)
	for i := 0; i < 3; i++ {
		filename, err = helper_os.GenUniqueFilename("/bin/ls", 10, func(name string) string {
			return name + "_" + helper_gen.UuidShort()
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(filename)
	}
}
