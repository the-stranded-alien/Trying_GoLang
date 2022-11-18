package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("GO Routine:\t", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GO Routine:\t", runtime.NumGoroutine())
	}

	wg.Wait()

	fmt.Println("CPU:\t\t", runtime.NumCPU())
	fmt.Println("GO Routine:\t", runtime.NumGoroutine())

	fmt.Println("Count:", counter)

}
