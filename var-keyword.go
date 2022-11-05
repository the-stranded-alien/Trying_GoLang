package main

import "fmt"

// Declare the variable "y"
// Assign the value 43
// declare and assign = initialization
var y = 43

// Declare there is a variable with identifier "z"
// and that the variable with the identifier "z" is of TYPE int
// assigns the ZERO value of type INT to "z"
// false for booleans, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps
var z int

func main() {
	// short declaration operator
	// DECLARE A VARIABLE and ASSIGN a VALUE at Same Time (of a certain type)
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	foobar()
	fmt.Println(z)
}

func foobar() {
	fmt.Println("again: ", y)
}
