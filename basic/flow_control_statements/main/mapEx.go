package main

import "fmt"

func main() {
	testMap()
	println("==================")
	fmt.Println(mapTest1)
	println("==================")
	fmt.Println(mapTest2)
	println("==================")
	actionOnMap()
	println("==================")
}

/*
MAP(хеш-таблица/ассоциативный массив) - не гарантирует порядок элементов
Нулевое значение карты равно nil. На nil карте нет ключей, и ключи нельзя добавить - особенность
var m = map[int](- это ключ)string(- это значение){
        1: "A",
        2: "B", - запятая в конце обязательна!!!
}
m[2] - получение значения из мапы
m[2] = "С" - добавление или перезапись значения
m[3] - даст пустую строку (так как ключа нет)

map:= make(map[int]string) - еще один способ создания мапы

test, isExist := m[3] - присвоит значение из мапы в test, isExist = будет true/false  в зависимости если такой ключ в мапе или нет
delete(m, 1) - удаление пары ключ+значение
*/

func testMap() {
	var m = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// Добавление структуры в значение карты
type VertexTest struct {
	Lat, Long float64
}

var mapTest1 = map[string]VertexTest{
	"Bell Labs": VertexTest{
		40.68433, -74.39967,
	},
	"Google": VertexTest{
		37.42202, -122.08408,
	},
}

var mapTest2 = map[string]VertexTest{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

// Добавление, изменение, удаление, получение
func actionOnMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
