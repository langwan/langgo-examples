package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core"
)

func main() {
	langgo.SetEnvName(core.Development)
	langgo.Run()
	fmt.Printf("mode = %s\n", langgo.GetEnvName())
}
