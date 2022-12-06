package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/mysql"
)

func main() {
	langgo.Run(&mysql.Instance{})
	var i int64
	mysql.Main().Raw("SELECT 1").Scan(&i)
	fmt.Println(i)
	mysql.Get("slave").Raw("SELECT 2").Scan(&i)
	fmt.Println(i)
}
