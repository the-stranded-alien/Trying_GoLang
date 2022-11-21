package main

import (
	"fmt"
	"log"
	"math"
)

type mathError struct {
	lat  string
	long string
	err  error
}

func (n mathError) Error() string {
	return fmt.Sprintf("a math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrtFun(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrtFun(f float64) (float64, error) {
	if f < 0 {
		me := fmt.Errorf("math redux: square root of negative number")
		return 0, mathError{"50.2289 N", "99.4656 W", me}
	}
	return math.Sqrt(f), nil
}
