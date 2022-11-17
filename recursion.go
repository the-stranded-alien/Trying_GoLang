package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println(n)

	n2 := loopFac(4)
	fmt.Println(n2)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func loopFac(num int) int {
	total := 1
	for ; num > 0; num-- {
		total *= num
	}
	return total
}
