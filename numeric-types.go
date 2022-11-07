package main

import "fmt"

var i int
var f float64

func main() {
	i = 42
	f = 42.34534
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
