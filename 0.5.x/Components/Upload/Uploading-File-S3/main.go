package main

import (
	"context"
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/upload"
	helper_s3 "github.com/langwan/langgo/helpers/s3"
	"time"

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
	s3Client, err := helper_s3.NewClient("xxxxxx", "xxxxxx", "xxxxxx", "xxxxxx", "xxxxxx", helper_s3.WithTimeout(time.Hour, time.Hour, time.Hour))
	if err != nil {
		panic(err)
	}
	s3Writer := upload.S3Writer{
		S3Client: s3Client,
	}

	partList := MyPartList{}

	upload.UploadFile(context.Background(), "../../../samples/sample.mov", "upload_test.mov", &partList, &s3Writer, &Listener{})
}
