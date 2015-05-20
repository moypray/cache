package main

import (
	"fmt"
	"github.com/HuKeping/cache"
)

func Print(item cache.Item, buffer *cache.LLRB) bool {
	i, ok := item.(cache.Int)
	if !ok {
		return false
	}
	fmt.Println(int(i))
	return true
}

func main() {
	buffer := cache.New()
	buffer.ReplaceOrInsert(cache.Int(1))
	buffer.ReplaceOrInsert(cache.Int(2))
	buffer.ReplaceOrInsert(cache.Int(3))
	buffer.ReplaceOrInsert(cache.Int(4))
	buffer.DeleteMin()
	buffer.Delete(cache.Int(4))
	buffer.AscendGreaterOrEqual(buffer.Min(), Print)
}
