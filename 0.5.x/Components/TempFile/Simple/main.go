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
	defer tempfile.Get().RemoveFile(filename)
	data, err := tempfile.Get().ReadFile(filename, false)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(data))

}
