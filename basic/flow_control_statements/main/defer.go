package main

import "fmt"

/*
Defer: выполнение перед выходом из функции (до return)
всегда выполниться
Когда используется:
- несколько вызовов необходимо сделать
- освобождение ресурсов (закрыть файл как пример)
*/
func main() {
	defer func() {
		fmt.Println("-")
	}() //вызов анонимной функции
	fmt.Println("+")
	example() //В консоли: + 2 1 -
}

func test() {
	fmt.Println("1")
}

func example() {
	defer test()
	fmt.Println("2")
}

/*
Отложенные вызовы функций помещаются в стек. (LIFO - последним пришел — первым вышел)
Пример:
func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
} В консоли будет вывод: от 9 до 0
*/
