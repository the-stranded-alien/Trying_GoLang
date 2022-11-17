package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)
	x++
	{
		y := 42
		fmt.Println(y)
	}
	//fmt.Println(y)
	fmt.Println(x)

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
