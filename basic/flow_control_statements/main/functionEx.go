package main

import (
	"fmt"
	"math"
)

/*
Функции тоже являются значениями. Их можно передавать так же, как и другие значения.
Значения функции могут использоваться в качестве аргументов функции и возвращаемых значений.
*/

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
