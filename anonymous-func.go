package main

import "fmt"

func main() {
	foofunc()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("The value of x is:", x)
	}(42)
}

func foofunc() {
	fmt.Println("Func ran")
}
