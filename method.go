package main

import "fmt"

type person2 struct {
	fn string
	ln string
}

type secretAgent2 struct {
	person2
	ltk bool
}

func (s secretAgent2) speak() {
	fmt.Println("I am,", s.fn, s.ln)
}

func main() {
	sa1 := secretAgent2{
		person2: person2{
			"James",
			"Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
}
