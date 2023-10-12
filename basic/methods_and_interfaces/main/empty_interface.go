package main

import "fmt"

// Тип интерфейса, в котором не указаны методы, известен как пустой интерфейс
// Такой интерфейс моет содержать значение любого типа
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
