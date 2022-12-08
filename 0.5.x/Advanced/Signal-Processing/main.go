package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"os"
	"syscall"
)

func main() {
	langgo.Run()
	done := make(chan struct{})
	langgo.SignalHandlers(func(sig os.Signal) {
		fmt.Printf("receive sig is %s\n", sig.String())
	}, syscall.SIGHUP, syscall.SIGUSR1)
	langgo.SignalNotify()
	fmt.Printf("pid = %d\n", os.Getpid())
	<-done
}
