package main

import "fmt"

func main() {
	s1 := "Hello, world!"
	s2 := `Hello,
			World!`
	fmt.Println(s1)
	s1 = "hello"
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)
	fmt.Println(s2)

	bs1 := []byte(s1)
	fmt.Println(bs1)
	fmt.Printf("%T\n", bs1)

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U", s1[i])
	}
	fmt.Println("")
	for i, v := range s1 {
		fmt.Printf("at index position %d we have hex %#x \n", i, v)
	}

}
