package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/hello"
)

func main() {
	langgo.Run(&hello.Instance{Message: "hello langgo"})
	langgo.Logger("app", "hello").Info().Str("message", hello.Get().Message).Send()
}
