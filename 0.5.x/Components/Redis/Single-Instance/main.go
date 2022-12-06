package main

import (
	"fmt"
	"github.com/langwan/langgo"
	"github.com/langwan/langgo/components/redis"
	"time"
)

func main() {
	langgo.Run(&redis.Instance{})
	redis.Main().Set("key", "langgo", 10*time.Second)
	val := redis.Main().Get("key")
	fmt.Println(val)
}
