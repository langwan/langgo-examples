package main

import (
	"context"
	"fmt"
	helper_command "github.com/langwan/langgo/helpers/command"
)

func main() {
	output, err := helper_command.Execute(context.Background(), "/Users/langwan/go/go1.19.1/bin/go", []string{"help"}...)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
