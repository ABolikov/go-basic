package main

import "fmt"

/*
Struct (структуры) - объекты - характеристики:
Наименование функции конструктор начинается со слова New (так принято делать)
Встраивание структур:

type Vehicle struct {
Brand string
Model string
Goe Geo - встраивание с отдельным полем
Пример в: basic/flow_control_statements/main/struct2.go
или
Geo - встраивание без отдельного поля (как будто поля из Geo struct находятся в Vehicle struct) аналогично с интерфейсами
Пример в: basic/flow_control_statements/main/struct1.go
}

type Geo struct {
Latitude string
Longitude string
}

Хранения ссылки на структуру того же типа:  basic/flow_control_statements/main/struct3.go
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	p := &v
	p.X = 55 //= записи (*p).X = 55
	(*p).Y = 10
	fmt.Println(v)
}
