package main

import (
	"fmt"
	helper_lru "github.com/langwan/langgo/helpers/lru"
)

func main() {
	lru := helper_lru.New[string, int](2)
	lru.Set("x", 1)
	lru.Set("y", 2)
	if val, ok := lru.Get("x"); ok {
		fmt.Println("x", val)
	}
	lru.Set("z", 3)
	if val, ok := lru.Get("y"); ok {
		fmt.Println("y", val)
	}
}
