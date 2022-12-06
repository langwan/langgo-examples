package main

import (
	"fmt"
	helper_os "github.com/langwan/langgo/helpers/os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			id := helper_os.GetGoroutineId()
			fmt.Printf("goroutine %d\n", id)
			wg.Done()
		}()
	}
	wg.Wait()
}
