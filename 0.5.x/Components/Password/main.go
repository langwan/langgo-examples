package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/password"
)

func main() {
	langgo.Run(&password.Instance{})
	hash := password.Hash("123456")
	fmt.Println(hash)
}
