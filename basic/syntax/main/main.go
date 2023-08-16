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
        . "github.com/vigo5190/goimports-example/b" - импорт с доступом ко всему имеющемуся в пакете в текущем без указания имени пакета
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
