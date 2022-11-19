package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go sendSelect(even, odd, quit)

	// receive
	receiveSelect(even, odd, quit)

	fmt.Println("About to exit")
}

func receiveSelect(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel:", v)
		case v := <-o:
			fmt.Println("From the odd  channel:", v)
		case v := <-q:
			fmt.Println("From the quit channel:", v)
			return
		}
	}
}

func sendSelect(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// close(e)
	// close(o)
	q <- 0
}
