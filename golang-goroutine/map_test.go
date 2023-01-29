package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go AddToMap(data, i, group)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Done")

}