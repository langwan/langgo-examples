package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/sqlite"
)

func main() {
	langgo.Run(&sqlite.Instance{Path: "./database.db"})
	var i int64
	sqlite.Get().Raw("SELECT 1").Scan(&i)
	fmt.Println(i)
}
