package main

import "fmt"

func main() {
	fmt.Println("Loops")
	//for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	for i := 0; i <= 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Outer loop: %d\t & Inner Loop: %d\n", i, j)
		}
	}
	h := 1
	for h < 10 {
		fmt.Println(h)
		h++
	}

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}
