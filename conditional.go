package main

import "fmt"

func main() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}

	if x := 42; x == 42 {
		fmt.Println("here")
	}
	y := 7
	if y == 6 {
		fmt.Println("y = 6")
	} else if y == 7 {
		fmt.Println("y = 7")
	} else {
		fmt.Println("y != 6 && y != 7")
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
