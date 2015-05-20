package main

import (
	"fmt"
	"github.com/HuKeping/cache"
	"time"
)

type Var struct {
	Expiry time.Time `json:"expiry,omitempty"`
	ID     string    `json:"id",omitempty`
}

// We will order the node by `Time`
func (x Var) Less(than cache.Item) bool {
	return x.Expiry.Before(than.(Var).Expiry)
}

func PrintVar(item cache.Item, buffer *cache.LLRB) bool {
	i, ok := item.(Var)
	if !ok {
		return false
	}

	if i.Expiry.Before(time.Now()) {
		fmt.Printf("Item expired and to be deleted : %v\n", i)
		buffer.Delete(i)
	}
	return true
}

func main() {
	buffer := cache.New()
	var1 := Var{
		Expiry: time.Now().Add(time.Second * 10),
		ID:     "var5",
	}
	var2 := Var{
		Expiry: time.Now().Add(time.Second * 20),
		ID:     "var4",
	}
	var3 := Var{
		Expiry: time.Now().Add(time.Second * 30),
		ID:     "var3",
	}
	var4 := Var{
		Expiry: time.Now().Add(time.Second * 40),
		ID:     "var2",
	}
	var5 := Var{
		Expiry: time.Now().Add(time.Second * 50),
		ID:     "var1",
	}

	buffer.ReplaceOrInsert(var1)
	buffer.ReplaceOrInsert(var2)
	buffer.ReplaceOrInsert(var3)
	buffer.ReplaceOrInsert(var4)
	buffer.ReplaceOrInsert(var5)

	go func() {
		for {
			time.Sleep(time.Second * 10)
			fmt.Printf("[VAR] Refreshing begin ...\n")
			buffer.AscendGreaterOrEqual(buffer.Min(), PrintVar)
			fmt.Printf("[VAR] Refreshing end...\n\n")
		}
	}()

	time.Sleep(time.Minute)
}
