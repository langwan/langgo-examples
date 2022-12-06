package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/tempfile"
)

func main() {
	langgo.Run(&tempfile.Instance{Base: "tmp"})
	filename, err := tempfile.Get().CreateFile([]byte("langgo"), 0644)
	if err != nil {
		panic(err)
		return
	}

	data, err := tempfile.Get().ReadFile(filename, true)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(data))

}
