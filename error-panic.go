package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer def2()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("We are here!")
}

func def2() {
	fmt.Println("When log.Panicln() is called, deferred functions run")
}
