package main

import (
	"fmt"
	helper_gen "github.com/langwan/langgo/helpers/gen"
)

func main() {
	uuid := helper_gen.Uuid()
	fmt.Println(uuid)
	uuid = helper_gen.UuidNoSeparator()
	fmt.Println(uuid)
	uuid = helper_gen.UuidShort()
	fmt.Println(uuid)
}
