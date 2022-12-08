package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/s3"
)

func main() {
	langgo.Run(&s3.Instance{})
	s3Client, err := s3.Get().NewClient(&s3.Client{
		Endpoint:        "oss-cn-hangzhou.aliyuncs.com",
		AccessKeyId:     "LTAI5tRiZ2Djq1qFBaBkfofr",
		AccessKeySecret: "hP2If3OJT2erbfUzjurppDi08oviWY",
		BucketName:      "banyun-files",
		Region:          "oss-cn-hangzhou",
	})
	if err != nil {
		panic(err)
	}
	head, err := s3Client.HeadObject("Homework - 1028.mp4")
	if err != nil {
		return
	}
	fmt.Println(head)
}
