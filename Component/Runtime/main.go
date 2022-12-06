package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
	"github.com/langwan/langgo/components/sqlite"
)

func main() {
	langgo.Run(&sqlite.Instance{Path: "./database.db"})
	langgo.LoadComponents(&hello.Instance{Message: "hello langgo"})
	fmt.Println(hello.Get().Message)
}
