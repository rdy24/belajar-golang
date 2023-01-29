package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := &sync.Pool{}

	pool.New = func() interface{} {
		return "New"
	}

	pool.Put("Tes")
	pool.Put("Data")
	pool.Put("Pool")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}