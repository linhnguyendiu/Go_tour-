package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i < 100; i++ {
		oldZ := z
		z -= (z*z - x) / (2 * z)
		delta := math.Abs(z - oldZ)
		fmt.Printf("Iteration %d: z = %v\n", i, z)
		if delta < 1e-10 {
			break
		}
	}
	return z
}

func main() {
	x := 3.0
	sqrt := Sqrt(x)
	fmt.Printf("Sqrt(%v) = %v\n", x, sqrt)
	fmt.Printf("math.Sqrt(%v) = %v\n", x, math.Sqrt(x))
	fmt.Printf("Delta(%v) = %v\n", x, sqrt-math.Sqrt(x))
}
