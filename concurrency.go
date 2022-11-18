package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GO Routines\t", runtime.NumGoroutine())

	wg.Add(1)
	go loop1()
	loop2()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GO Routines\t", runtime.NumGoroutine())
	wg.Wait()
}

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Println("l1:", i)
	}
	wg.Done()
}

func loop2() {
	for i := 0; i < 10; i++ {
		fmt.Println("l2:", i)
	}
}
