package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/ffmpeg"
)

func main() {
	langgo.Run(&ffmpeg.Instance{})
	duration, err := ffmpeg.Get().Duration("../../../samples/sample.mov")
	if err != nil {
		return
	}
	fmt.Println(duration)
}
