package main

import "fmt"

func main() {
	ev := make(chan int)
	od := make(chan int)
	qu := make(chan bool)

	// send
	go sendOk(ev, od, qu)

	// receive
	receiveOk(ev, od, qu)

	d := make(chan int)
	go func() {
		d <- 66
		close(d)
	}()

	v, ok := <-d
	fmt.Println(v, ok)

	v, ok = <-d
	fmt.Println(v, ok)

	fmt.Println("About to exit")
}

func receiveOk(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel:", v)
		case v := <-o:
			fmt.Println("From the odd  channel:", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from comma ok", i, ok)
				return
			} else {
				fmt.Println("from comma ok ", i)
			}
		}
	}
}

func sendOk(e, o chan<- int, q chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}
