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
		ID:     "var1",
	}
	var2 := Var{
		Expiry: time.Now().Add(time.Second * 20),
		ID:     "var2",
	}
	var3 := Var{
		Expiry: time.Now().Add(time.Second * 30),
		ID:     "var3",
	}
	var4 := Var{
		Expiry: time.Now().Add(time.Second * 40),
		ID:     "var4",
	}
	var5 := Var{
		Expiry: time.Now().Add(time.Second * 50),
		ID:     "var5",
	}

	buffer.Add(var1, cache.MOD_REPLACE)
	buffer.Add(var2, cache.MOD_REPLACE)
	buffer.Add(var3, cache.MOD_REPLACE)
	buffer.Add(var4, cache.MOD_REPLACE)
	buffer.Add(var5, cache.MOD_REPLACE)

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
