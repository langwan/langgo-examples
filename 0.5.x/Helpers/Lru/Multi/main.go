package main

import (
	"fmt"
	helper_lru "github.com/langwan/langgo/helpers/lru"
)

func main() {
	lruAccount := helper_lru.New[string, int](2)
	lruAccount.Set("x", 1)
	lruAccount.Set("y", 2)
	lruAccount.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true
	})
	lruOrder := helper_lru.New[int, int](2)
	lruOrder.Set(1, 1)
	lruOrder.Set(2, 2)
	lruOrder.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true
	})
}
