package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello, world", 42, true)
	fmt.Println(n)
	fmt.Println(err)
	//n, _ := fmt.Println("Hello, world", 42, true)
	//fmt.Println(n)
}
