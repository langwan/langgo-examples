package main

import (
	"context"
	"fmt"
	helper_command "github.com/langwan/langgo/helpers/command"
)

func main() {
	output, err := helper_command.Execute(context.Background(), "ls", []string{"-la"}...)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
