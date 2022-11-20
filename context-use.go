package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Error Check 1:", ctx.Err())
	fmt.Println("Number of GO Routines 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Error Check 2:", ctx.Err())
	fmt.Println("Number of GO Routines 2:", runtime.NumGoroutine())

	fmt.Println("About to Cancel Context")
	cancel()
	fmt.Println("Cancelled Context")

	time.Sleep(time.Second * 2)
	fmt.Println("Error Check 3:", ctx.Err())
	fmt.Println("Number of GO Routines 3:", runtime.NumGoroutine())
}
