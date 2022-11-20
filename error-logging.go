package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Error Happened", err)
		log.Println("Error Happened", err)
		log.Fatalln(err)
	}
}
