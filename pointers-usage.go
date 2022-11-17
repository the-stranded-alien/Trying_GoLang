package main

import "fmt"

// EVERYTHING IN GOLANG IS "PASS BY VALUE"

func main() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	point(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func point(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
