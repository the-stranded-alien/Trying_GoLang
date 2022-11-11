package main

import "fmt"

const (
	c1 string  = "hello"
	c2 int     = 23
	c3 float64 = 24.5
)

func main() {
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)
	fmt.Printf("%T\n", c3)

}
