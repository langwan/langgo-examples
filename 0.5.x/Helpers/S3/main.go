package main

import (
	"fmt"
	helper_s3 "github.com/langwan/langgo/helpers/s3"
	"time"
)

func main() {
	client, err := helper_s3.NewClient("oss-cn-hangzhou.aliyuncs.com", "LTAI5tRiZ2Djq1qFBaBkfofr", "hP2If3OJT2erbfUzjurppDi08oviWY", "banyun-files", "oss-cn-hangzhou", helper_s3.WithTimeout(time.Hour, time.Hour, time.Hour))
	if err != nil {
		panic(err)
	}
	head, err := client.HeadObject("Homework - 1028.mp4")
	if err != nil {
		return
	}
	fmt.Println(head)
}
