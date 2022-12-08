package main

import (
	"fmt"
	helper_s3 "github.com/langwan/langgo/helpers/s3"
	"time"
)

func main() {
	client, err := helper_s3.NewClient("xxxxxx", "xxxxxx", "xxxxxx", "xxxxxx", "xxxxxx", helper_s3.WithTimeout(time.Hour, time.Hour, time.Hour))
	if err != nil {
		panic(err)
	}
	head, err := client.HeadObject("Homework - 1028.mp4")
	if err != nil {
		return
	}
	fmt.Println(head)
}
