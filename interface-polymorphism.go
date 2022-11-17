package main

import "fmt"

type person3 struct {
	fn string
	ln string
}

type secretAgent3 struct {
	person3
	ltk bool
}

func (s secretAgent3) spk() {
	fmt.Println("I am", s.fn, s.ln, " - The secret agent speak")
}

func (p person3) spk() {
	fmt.Println("I am", p.fn, p.ln, " - The person speak")
}

type human interface {
	spk()
}

func hum(h human) {
	fmt.Println("I was passed into hum", h)
}

func hum2(h human) {
	switch h.(type) {
	case person3:
		fmt.Println("Am just a person", h.(person3).fn)
	case secretAgent3:
		fmt.Println("Am a secret agent", h.(secretAgent3).fn)
	}
	fmt.Println("I was passed into hum2", h)
}

type hotdogs int

func main() {
	sa1 := secretAgent3{
		person3: person3{
			"James",
			"Bond",
		},
		ltk: true,
	}
	p1 := person3{
		fn: "Dr.",
		ln: "Yes",
	}
	fmt.Println(sa1)
	sa1.spk()
	fmt.Println(p1)
	p1.spk()
	hum(sa1)
	hum(p1)

	hum2(sa1)
	hum2(p1)

	// conversion
	var xx hotdogs = 42
	fmt.Println(xx)
	fmt.Printf("%T\n", xx)
	var yy int
	yy = int(xx)
	fmt.Println(yy)
	fmt.Printf("%T\n", yy)
}
