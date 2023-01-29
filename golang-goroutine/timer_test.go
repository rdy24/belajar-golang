package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5 * time.Second, func(){
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)
	fmt.Println(time.Now())

	go func ()  {
		time.Sleep(10 * time.Second)
		ticker.Stop()
	} ()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(5 * time.Second)
	fmt.Println(time.Now())

	
	for time := range channel {
		fmt.Println(time)
	}
}