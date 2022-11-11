package main

import "fmt"

const (
	i1 = iota
	i2
	i3
)

const (
	i4 = iota
	i5
	i6
)

func main() {
	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)
	fmt.Println(i4)
	fmt.Println(i5)
	fmt.Println(i6)

}
