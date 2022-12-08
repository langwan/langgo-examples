package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/core/log"
	"os"
	"syscall"
	"time"
)

func main() {
	langgo.Run()
	log.SetCuttingSignal(syscall.SIGUSR1)
	fmt.Printf("pid = %d\n", os.Getpid())
	langgo.SignalNotify()
	for {
		log.Logger("app", "main").Info().Timestamp().Send()
		time.Sleep(time.Second)
	}
}
