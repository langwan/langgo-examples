package main

import (
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/ffmpeg"
)

func main() {
	langgo.Run(&ffmpeg.Instance{})
	ffmpeg.Get().Transcoding("../../../samples/sample.mov", "./output.mp4", true)
}
