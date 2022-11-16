package main

import "fmt"

func main() {
	defer foofoo()
	barbar()
}

func foofoo() {
	fmt.Println("foo")
}

func barbar() {
	fmt.Println("bar")
}
