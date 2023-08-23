/*
https://metanit.com/go/tutorial/2.3.php
*/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
uint - безнаковое число (только положительное)
int - число со знаком (могу быть отрицательные) - int8,16,32,64 - количество занимаемы бит
float - числа c плавающей точкой
*/
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	cast()
}

func cast() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
