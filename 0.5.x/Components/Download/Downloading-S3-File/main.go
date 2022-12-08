package main

import (
	"context"
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/download"
	"github.com/langwan/langgo/components/s3"
	helper_progress "github.com/langwan/langgo/helpers/progress"
)

type MyListener struct {
}

func (l MyListener) ProgressChanged(event *helper_progress.ProgressEvent) {
	fmt.Println(event)
}

func main() {
	langgo.Run(&download.Instance{}, &s3.Instance{})
	s3Client, err := s3.Get().NewClient(&s3.Client{
		Endpoint:        "xxxxxx",
		AccessKeyId:     "xxxxxx",
		AccessKeySecret: "xxxxxx",
		BucketName:      "xxxxxx",
		Domain:          "",
		Region:          "xxxxxxu",
	})
	if err != nil {
		panic(err)
	}
	reader := download.S3Reader{
		ObjectName: "Homework - 1028.mp4",
		Client:     s3Client,
	}
	download.Download(context.Background(), "./output.mp4", &reader, &MyListener{})
}
