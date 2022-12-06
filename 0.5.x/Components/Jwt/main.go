package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/jwt"
)

func main() {
	langgo.Run(&jwt.Instance{Secret: "123456"})
	payload := struct {
		Name string
	}{Name: "langwan"}
	sign, err := jwt.Sign(payload)
	if err != nil {
		panic(err)
	}
	err = jwt.Verify(sign)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("verify ok")
	}
}
