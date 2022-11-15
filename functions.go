package main

import "fmt"

func main() {
	foo2()
	bar("James")
	rs1 := woo("Jimmy")
	fmt.Println(rs1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
	sv := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sv)
}

// func (r receiver) identifier(parameters) (return(s)) {...}

func foo2() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sumval := 0
	for _, v := range x {
		sumval += v
	}
	return sumval
}
