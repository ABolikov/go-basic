package main

import "fmt"

//Хранения ссылки на структуру того же типа
// Структура не может иметь поле, которое представляет тип этой же структуры
//type node struct{
//	value int
//	next node
//}

type node struct {
	value int
	next  *node //<- через указатель структуры
}

// рекурсивный вывод списка
func printNodeValue(n *node) {

	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}
func main() {

	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	fmt.Println("============")
	printNodeValue(&first)
}
