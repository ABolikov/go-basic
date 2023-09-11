package main

import (
	"fmt"
)

func fibonacci() func() int {
	x := 0
	y := 0
	sum := 0
	return func() int {
		sum = x + y
		if sum == 0 {
			x += 1
			return y
		}
		if sum == 1 && y == 0 {
			x = 0
			y = 1
			return y
		} else {
			x = y
			y = sum
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
