package main

import "fmt"

const (
	_ = iota
	// kb = 1024
	kb = 1 << (iota * 10)
	// mb = 1024 * kb
	mb = 1 << (iota * 10)
	// gb = 1024 * mb
	gb = 1 << (iota * 10)
)

func main() {

	b1 := 4
	fmt.Printf("%d\t\t%b\n", b1, b1)

	b2 := b1 << 1
	fmt.Printf("%d\t\t%b\n", b2, b2)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
