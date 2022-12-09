package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
	"github.com/langwan/langgo/components/sqlite"
)

func main() {
	langgo.Run(&hello.Instance{Message: "hello langgo"}, &sqlite.Instance{Path: "./database.db"})
	langgo.Logger("app", "hello").Info().Str("message", hello.Get().Message).Send()
	var i int64
	sqlite.Get().Raw("SELECT 1").Scan(&i)
	langgo.Logger("db", "main").Debug().Int64("i", i).Send()
}
