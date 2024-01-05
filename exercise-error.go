package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrtx(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	x := float64(16)
	result, err := Sqrtx(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Sqrt(%v) = %v", x, result)
	}
}
