package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("GO Routine:\t", runtime.NumGoroutine())

	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("GO Routine:\t", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("GO Routine:\t", runtime.NumGoroutine())

	fmt.Println("Count:", counter)

}
