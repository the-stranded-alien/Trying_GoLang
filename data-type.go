package main

import "fmt"

var a = 42

// Declared the Variable with the identifier "b" is of type STRING and assigned the value "hello"
var b = "hello"

// GO is a STATIC programming language, a variable is declared to hold a VALUE of a certain type

func main() {
	fmt.Println("GO data types")
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// Can't Do
	//b = 43
	//fmt.Printf("%T\n", a)
	//fmt.Println(b)

}
