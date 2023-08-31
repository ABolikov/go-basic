package main

import (
	"fmt"
	"math"
	"math/rand"
)

/*
 Документация:
https://pkg.go.dev/fmt
 https://pkg.go.dev/math
*/

/*
import (
        "github.com/vigo5190/goimports-example/a" - обычный импорт пакета
        foo "github.com/vigo5190/goimports-example/a" - импорт c синонимом
        . "github.com/vigo5190/goimports-example/b" - импорт с доступом ко всему имеющемуся в пакете, в текущем - без указания имени пакета
        _ "github.com/vigo5190/goimports-example/c" - импорт для выполнения функции init() из пакета
    )
*/

var c, python, java bool

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	fmt.Println(c, python, java)
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(k, c, python, java)

}

// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

/*
Именованный возврат (хорош для коротких функций)
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//Константы

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func example() {
	fmt.Printf("v is of type %T\n", needFloat(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big)) -> выдаст ошибку, так как инт не держит такой объем
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}
