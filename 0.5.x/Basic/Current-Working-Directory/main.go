package main

import (
	"fmt"
	"github.com/langwan/langgo"
)

func main() {
	langgo.SetWorkDir("./bin")
	workDir := langgo.GetWorkDir()
	fmt.Println(workDir)
}
