package main

import "fmt"

func main() {
	c := make(chan int)
	// c <- 42 // Won't work
	// This would work
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
	// This would also work, buffer
	d := make(chan int, 1)
	d <- 42
	fmt.Println(<-d)
	// Unsuccessful buffer
	//e := make(chan int, 1)
	//e <- 43
	//e <- 44
	//fmt.Println(<-e)
}
