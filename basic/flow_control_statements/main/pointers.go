package main

import "fmt"

/*
Указатели (на адрес ячейки памяти):
& - оператор
* - оператор
Тип *T является указателем на T значение

Получение указателя на тип данных string:

phrase := "test"
phrasePtr := &phrase - получает адрес в памяти (использовать нужно когда точно везде надо изменить объект)
fmt.Println(phrasePtr)
fmt.Println(*phrasePtr) - выводит значение из указанного адреса
*/

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21 // установка значения по указателю
	fmt.Println(i)

	p = &j
	*p = *p / 37 // деление j через указатель
	fmt.Println(j)

	fmt.Println(*p)
}
