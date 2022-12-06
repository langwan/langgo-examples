package main

import (
	"fmt"
	helper_gen "github.com/langwan/langgo/helpers/gen"
)

func main() {
	rs, err := helper_gen.RandString(10)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs)
}
