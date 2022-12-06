package main

import (
	"fmt"
	helper_gen "github.com/langwan/langgo/helpers/gen"
)

func main() {
	for i := 0; i < 10; i++ {
		ri := helper_gen.RandInt(10, 100)
		fmt.Printf("%d ", ri)
	}
}
