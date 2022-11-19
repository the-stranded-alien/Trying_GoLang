package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	// go sendRange(c)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // necessary
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("About to exit")
}

func sendRange(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
