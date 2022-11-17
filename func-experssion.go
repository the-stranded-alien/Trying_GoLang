package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("my first func expression")
	}

	f2 := func(x int) {
		fmt.Println("Value of x is:", x)
	}

	f()
	f2(50)
}
