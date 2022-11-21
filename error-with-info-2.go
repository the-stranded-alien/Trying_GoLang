package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	_, err := sqrtNew(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrtNew(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("math error again: square root of negative number: %v", f)
	}
	return math.Sqrt(f), nil
}
