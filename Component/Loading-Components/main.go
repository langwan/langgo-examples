package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
	"github.com/langwan/langgo/components/sqlite"
)

func main() {
	langgo.Run(&hello.Instance{}, &sqlite.Instance{})
	fmt.Println(hello.Get().Message)
	var i int64
	sqlite.Get().Raw("SELECT 1").Scan(&i)
	fmt.Println(i)
}
