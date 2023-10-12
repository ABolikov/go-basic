package main

/*
Программы Go выражают состояние ошибки с помощью error значений.

Тип error представляет собой встроенный интерфейс (аналогично fmt.Stringer):

type error interface {
    Error() string
}

Использование в коде:
Функции часто возвращают error значение, а вызывающий код должен обрабатывать ошибки, проверяя, равна ли ошибка nil.

Пример:
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)

*/

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
