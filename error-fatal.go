package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer def()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("We are here!")
}

func def() {
	fmt.Println("When os.Exit()/log.Fatalln() is called, deferred functions don't run")
}
