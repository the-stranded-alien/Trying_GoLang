package main

import "fmt"

func main() {
	//sw := hello()
	//fmt.Println(sw)

	x := func() int {
		return 345
	}
	fmt.Printf("%T\n", x)

	xx := fun()
	fmt.Printf("%T\n", xx)

	fmt.Println(xx())
	fmt.Println(fun()())

}

func hello() string {
	return "Hello world"
}

func fun() func() int {
	return func() int {
		return 451
	}
}
