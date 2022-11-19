package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\n", c)

	uc := make(chan<- int, 2) // send
	uc <- 23
	uc <- 32
	fmt.Printf("%T\n", uc)
	// fmt.Println(<-uc) // Won't work

	uc2 := make(<-chan int, 2) // receive
	// uc2 <- 73
	// fmt.Println(<-uc2)
	fmt.Printf("%T\n", uc2)
}
