package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("We are here!")
}
