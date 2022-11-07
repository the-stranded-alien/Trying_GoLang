package main

import "fmt"

var bo bool

func main() {
	fmt.Println(bo)
	bo = true
	fmt.Println(bo)

	v1 := 7
	v2 := 42
	fmt.Println(v1 <= v2)
}
