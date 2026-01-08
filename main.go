package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const Eps = 1e-15

	prev, z := 0.0, 1.0

	for math.Abs(z-prev) > Eps {
		prev = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(24))
}
