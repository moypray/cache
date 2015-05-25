package main

import (
	"fmt"
	"github.com/HuKeping/cache"
)

func Print(item cache.Item, buffer *cache.LLRB) bool {
	i, ok := item.(cache.String)
	if !ok {
		return false
	}
	fmt.Println(cache.String(i))
	return true
}

func main() {
	buffer := cache.New()

	// For now it is not ordered.
	order := []cache.String{
		"ab", "aba", "abc", "a", "aa", "aaa", "b", "a-", "a!",
	}

	for _, i := range order {
		buffer.Add(cache.String(i), cache.MOD_REPLACE)
	}

	// We expect to get `ab`, `aba` and `abc`, otherwise something wrong happened.
	buffer.AscendRange(cache.String("ab"), cache.String("ac"), Print)
}
