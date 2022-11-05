package main

import "fmt"

var c int

type hotdog int

var x hotdog

func main() {
	c = 42
	fmt.Println(c)
	x = 78
	fmt.Println(x)
	//Can't do
	//c = x
	// CONVERTING NO CASTING
	c = int(x)
	fmt.Println(c)

}
