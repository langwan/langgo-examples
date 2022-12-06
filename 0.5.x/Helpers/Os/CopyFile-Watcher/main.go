package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
)

type Listener struct{}

func (l *Listener) ProgressChanged(event *helper_os.ProgressEvent) {
	fmt.Printf("event = %d, %d / %d\n", event.EventType, event.ConsumedBytes, event.TotalBytes)
}

func main() {
	_, err := helper_os.CopyFileWatcher("./ls.bk", "/bin/ls", nil, &Listener{})
	if err != nil {
		panic(err)
	}
}
