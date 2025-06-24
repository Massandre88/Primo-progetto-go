package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i, prev := 0, z; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if prev == z {
			return z
		}
		fmt.Println(x, z)
		prev = z
	}

	return z
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(Sqrt(float64(i)))
	}

}
