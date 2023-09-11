package main

import "fmt"

/*
Слайс/срез - гарантирует порядок элементов (динамический массив) смотрит на исходный массив: - длин (len) и вместимость (cap) - в него можно добавлять и удалять элементы
для создания слайса достаточно не указывать размер массива: animal := []string{}
week := [...] int{0,1,2,3,4,5,6,7}
weekend := week[6:8] (в указанную выборку попадает 6 и 7 индекс 8 - не включается) или week[6:] (до конца)
animals:= make([]string, 5, 10) - еще один способ создания
animals:= append(animals, "lion") - добавление в слайс, при добавлении, если исходный капасити мал - будет скопирован предыдущий массив, увеличится его капасити и будет добавлен новый элемент
*/

func main() {
	fmt.Println("======createSlice=======")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("Array: %d\n", primes)
	var s []int = primes[1:4]
	fmt.Printf("Slice: %d\n", s)

	fmt.Println("======slice=======")
	exampleSlice()

	fmt.Println("=======capacity=======")
	capacity()

	fmt.Println("======sliceMake=======")
	sliceMake()

	fmt.Println("======sliceAppend=======")
	sliceAppend()

	fmt.Println("======rangeSlice=======")
	indexAndValueForRangeSlice()

	fmt.Println("======rangeSliceTest=======")
	rangeSliceTest()
}

/*
Основное:
Срез не хранит никаких данных, он просто описывает раздел базового массива.
Изменение элементов среза изменяет соответствующие элементы базового массива.
Другие фрагменты, использующие тот же базовый массив, увидят эти изменения.
*/

func exampleSlice() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	//a := names[0:2]
	a := names[:2] //использование значения по умолчанию
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	c := names[1:]
	fmt.Println(c)
}

// Увеличить длину среза
func capacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	s = s[:4]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Создание слайсов с помощью make
// make(x, y, z) где: x-требуемый массив без указания размера, y - количество элементов массива, z - вместимость
func sliceMake() {
	a := make([]int, 5)
	printSlicesMake("a", a)

	b := make([]int, 0, 5)
	printSlicesMake("b", b)

	c := b[:2]
	printSlicesMake("c", c)

	d := c[2:5]
	printSlicesMake("d", d)
}

func printSlicesMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func sliceAppend() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

// range всегда возвращает два значения как и мапах, где первое - это индекс, второе - значение в этом индексе
var rangeSlice = []int{1, 2, 4, 8, 16, 32, 64, 128}

func indexAndValueForRangeSlice() {
	for i, v := range rangeSlice {
		fmt.Printf("index:%d = %d\n", i, v)
	}
}

// использование только одного из параметров при range
// Примеры записи:
//
//	 for i, _ := range pow
//		for _, value := range pow
//
// если нужен только индекс можно опустить вторую переменную: for i := range pow
func rangeSliceTest() {
	pow := make([]int, 10)
	//заполняем слайс
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	//выводим его значения
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
	//выводим индексы
	for i, _ := range pow {
		fmt.Printf("%d\n", i)
	}
}
