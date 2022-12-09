package main

import (
	"fmt"
	helper_lru "github.com/langwan/langgo/helpers/lru"
)

func main() {
	lru := helper_lru.New[string, int](2)
	lru.Set("x", 1)
	lru.Set("y", 2)
	lru.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true
	})
}
