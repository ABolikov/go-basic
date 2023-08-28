package main

import (
	"fmt"
	"math"
)

func main() {
	x := float64(2)
	fmt.Println("============")
	fmt.Println(Sqrt(x))
	fmt.Println("============")
	fmt.Println(math.Sqrt(x))
}

func Sqrt(x float64) float64 {
	z := float64(1)
	y := x
	for {
		z -= (z*z - x) / (2 * z)
		if y > z {
			y = z
		} else {
			break
		}
	}
	return y
}

func Sqrt2(x float64) float64 {
	z := float64(1)
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}
