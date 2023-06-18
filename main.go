package main

import (
	"fmt"
	"github.com/sunzhqr/HW/tree/main/hm/in_memory_cash/in_memory_cash"
)

func main() {
	cache := in_memory_cash.New()
	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
