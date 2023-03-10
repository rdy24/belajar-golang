package golanggoroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread: ", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total Go Routine: ", totalGoRoutine)
}