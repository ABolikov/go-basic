package main

import "fmt"

/*
Массивы (конкретное количество): - область памяти с множеством однотипных значений
var arr = [5]int{1,2,3,4,5} - объявление массива (по умолчанию ставиться 0 - значит их будет 5)
индексы с 0
animals:= make([]string, 5, 10) - еще один способ создания массива

Перебор массива - при этом можно за место индекса или значения использовать знак "_" - что бы исключить из обработки
arr := []string{"A","B","C"}
for index, letter:= range arr{
...////
}
range - значит перебираем

Перебор:
for condition {
///
}
condition -  условие для выхода из цикла
break, return -  для выхода из цикла
*/

func main() {
	fullFor()
	println("==================")
	notInitFor()
	println("==================")
	notPostFor()
	println("==================")
	array()
	println("==================")
}

// Полная запись [инициализация счетчика]; [условие]; [изменение счетчика] {}
func fullFor() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// Без инициализации = ; [условие]; [изменение счетчика] {}
func notInitFor() {
	sum := 0
	i := 0
	for ; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// Без инициализации  и счетчика = ; [условие]; {}
func notPostFor() {
	sum := 0
	k := 0
	for k < 5 {
		sum += k
		k++
	}
	fmt.Println(sum)

	//Краткая запись
	i := 1
	for i < 10 {
		fmt.Println(i * i)
		i++
	}
}

func array() {
	var users = [3]string{"Tom", "Alice", "Kate"}
	for index, value := range users {
		fmt.Println(index, value)
	}
	//Стандавртнвый варинт перебора, лен возращает длину массива
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
