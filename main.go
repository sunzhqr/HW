package main

import (
	"./in_memory_cash"
	"fmt"
)

func main() {
	Cache := in_memory_cash.Cache
	Cache.Set("userId", 42)
	userId := Cache.Get("userId")

	fmt.Println(userId)

	Cache.Delete("userId")
	userId = Cache.Get("userId")

	fmt.Println(userId)
}
