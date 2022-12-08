package main

import (
	"context"
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/download"
	helper_oss "github.com/langwan/langgo/helpers/oss"
	helper_progress "github.com/langwan/langgo/helpers/progress"
)

type MyListener struct {
}

func (l MyListener) ProgressChanged(event *helper_progress.ProgressEvent) {
	fmt.Println(event)
}

func main() {
	langgo.Run(&download.Instance{})
	ossClient, err := helper_oss.NewClient("xxxxxx", "xxxxxx", "xxxxxx", "xxxxxx")
	if err != nil {
		panic(err)
	}
	reader := download.OssReader{
		ObjectName: "Homework - 1028.mp4",
		Client:     ossClient,
	}
	download.Download(context.Background(), "./output.mp4", &reader, &MyListener{})
}
