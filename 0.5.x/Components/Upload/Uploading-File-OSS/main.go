package main

import (
	"context"
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/upload"
	helper_oss "github.com/langwan/langgo/helpers/oss"
	helper_progress "github.com/langwan/langgo/helpers/progress"
)

type MyPartList struct {
	uploadId string
	parts    []*upload.Part
}

func (m *MyPartList) RomoveParts() error {
	return nil
}

func (m *MyPartList) Load() ([]*upload.Part, error) {
	return nil, nil
}

func (m *MyPartList) SetList(parts []*upload.Part) {
	m.parts = parts
}

func (m *MyPartList) GetList() []*upload.Part {
	return m.parts
}

func (m *MyPartList) SavePart(part *upload.Part) error {
	return nil
}

func (m *MyPartList) GetUploadId() string {
	return m.uploadId
}

func (m *MyPartList) SetUploadId(uploadId string) error {
	m.uploadId = uploadId
	return nil
}

type Listener struct {
}

func (l Listener) ProgressChanged(event *helper_progress.ProgressEvent) {
	fmt.Println(event)
}

func main() {
	langgo.Run(&upload.Instance{})
	ossClient, err := helper_oss.NewClient("xxxxxx", "xxxxxx", "xxxxxx", "xxxxxx")
	if err != nil {
		panic(err)
	}
	writer := upload.OssWriter{
		Client: ossClient,
	}

	partList := MyPartList{}

	upload.UploadFile(context.Background(), "../../../samples/sample.mov", "upload_oss_test.mov", &partList, &writer, &Listener{})
}
