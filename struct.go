package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}
	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2)
	fmt.Println(p2.first, p2.last, p2.age)

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

	sa2 := secretAgent{
		person: p2,
		ltk:    false,
	}
	fmt.Println(sa2)

	pe := struct {
		first, last string
		age         int
	}{
		first: "James",
		last:  "Bond",
		age:   24,
	}
	fmt.Println(pe)
}
