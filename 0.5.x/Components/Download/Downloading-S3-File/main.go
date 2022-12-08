package main

import (
	"context"
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/download"
	helper_progress "github.com/langwan/langgo/helpers/progress"
	helper_s3 "github.com/langwan/langgo/helpers/s3"
	"time"
)

type MyListener struct {
}

func (l MyListener) ProgressChanged(event *helper_progress.ProgressEvent) {
	fmt.Println(event)
}

func main() {
	langgo.Run(&download.Instance{})
	s3Client, err := helper_s3.NewClient("xxxxxx", "xxxxxx", "xxxxxx", "xxxxxx", "xxxxxx", helper_s3.WithTimeout(time.Hour, time.Hour, time.Hour))
	if err != nil {
		panic(err)
	}
	reader := download.S3Reader{
		ObjectName: "Homework - 1028.mp4",
		Client:     s3Client,
	}
	download.Download(context.Background(), "./output.mp4", &reader, &MyListener{})
}
