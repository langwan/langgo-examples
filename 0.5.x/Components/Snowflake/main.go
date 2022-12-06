package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/snowflake"
)

func main() {
	langgo.Run(&snowflake.Instance{MachineID: 1})
	id := snowflake.Gen()
	fmt.Printf("Int64  ID: %d\n", id)
	id = snowflake.Gen()
	fmt.Printf("Int64  ID: %d\n", id)
	id = snowflake.Gen()
	fmt.Printf("Int64  ID: %d\n", id)
}
