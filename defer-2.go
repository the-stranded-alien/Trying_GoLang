package main

import "fmt"

func main() {
	var x int
	x++
	fmt.Println(x)
	i := cf()
	fmt.Println(i)
}

func cf() (i int) {
	defer func() { i++ }()
	return 1
}
