package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"langwan/langgo-examples/Advanced/Custom-Components/components/my"
)

func main() {
	langgo.Run(&my.Instance{})
	fmt.Println(my.Get().Message)
}
