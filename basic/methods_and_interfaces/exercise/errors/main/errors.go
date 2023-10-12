package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (float ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(float))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
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
	return y, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
