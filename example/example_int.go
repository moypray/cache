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
	buffer.Add(cache.Int(1), cache.MOD_NOREPLACE)
	buffer.Add(cache.Int(2), cache.MOD_NOREPLACE)
	buffer.Add(cache.Int(3), cache.MOD_NOREPLACE)
	buffer.Add(cache.Int(4), cache.MOD_NOREPLACE)
	buffer.DeleteMin()
	buffer.Delete(cache.Int(4))
	buffer.AscendGreaterOrEqual(buffer.Min(), Print)
}
