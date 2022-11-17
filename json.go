package main

import (
	"encoding/json"
	"fmt"
)

type person4 struct {
	First string
	Last  string
	Age   int
}

type person5 struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	p1 := person4{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person4{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}
	people := []person4{p1, p2}
	fmt.Println(people)

	// Marshal
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	// Unmarshal
	var people2 []person4
	err2 := json.Unmarshal(bs, &people2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(people2)

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Money","Age":27}]`
	bs2 := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs2)

	var people3 []person5
	err3 := json.Unmarshal(bs2, &people3)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(people3)

	for i, v := range people3 {
		fmt.Println("\nPerson Number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
