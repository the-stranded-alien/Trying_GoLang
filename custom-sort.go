package main

import (
	"fmt"
	"sort"
)

type per struct {
	first string
	age   int
}

func (p per) String() string {
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

// ByAge implements sort.Interface for []per based on the age field
type ByAge []per

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func main() {
	p1 := per{"James", 32}
	p2 := per{"Moneypenny", 27}
	p3 := per{"Q", 64}
	p4 := per{"M", 56}

	people := []per{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
