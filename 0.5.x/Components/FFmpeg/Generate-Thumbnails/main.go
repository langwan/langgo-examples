package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/ffmpeg"
	"time"
)

func main() {
	langgo.Run(&ffmpeg.Instance{})
	ffmpeg.Get().Thumbnail("../../../samples/sample.mov", "./output.jpg", time.Second, true)
}
