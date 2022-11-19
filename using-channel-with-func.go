package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go send(c)

	// receive
	//go receive(c)
	receive(c)

	fmt.Println("About to exit")
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}