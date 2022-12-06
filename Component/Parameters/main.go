package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
)

func main() {
	langgo.Run(&hello.Instance{Message: "hello langgo"})
	fmt.Println(hello.Get().Message)
}
