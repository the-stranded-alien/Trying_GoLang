package main

import (
	"fmt"
	"sync"
)

func main() {
	e := make(chan int)
	o := make(chan int)
	fanin := make(chan int)

	// send
	go sendFI(e, o)

	// receive
	go receiveFI(e, o, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

func receiveFI(e, o <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

func sendFI(e, o chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}
