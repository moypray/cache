# Buffer

## FYI

This project is based on GoLLRB which was written by [Petar Maymounkov](http://pdos.csail.mit.edu/~petar/).

Thanks for the great work of him and you can follow him on [Twitter @maymounkov](http://www.twitter.com/maymounkov)!

The origin project `GoLLRB` seems have slept for a long time and will no longer to be maintained, so I think
it's better to open a new repo and keep it fresh.

If anyone has any objections, please let me know.

## Example

A simple case for `int` items.

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
