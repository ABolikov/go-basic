package main

import (
	"fmt"
	"math"
)

/*
Интерфейсы представляют абстракцию поведения других типов.
Интерфейсы позволяют определять функции, которые не привязаны к конкретной реализации.
То есть интерфейсы определяют некоторый функционал, но не реализуют его.

Тип интерфейса определяется как набор сигнатур методов.

Значение типа интерфейса может содержать любое значение, реализующее эти методы.
Значение интерфейса -> значение определенного базового конкретного типа.
Вызов метода для значения интерфейса выполняет одноименный метод для его базового типа.

При этом выполняется не явная реализация интерфейса: нет никаких слов и указания на то что метод реализует какой интерфейс

Примечание:
Если значение самого интерфейса/структуры которая его реализует равно нулю(nil) то метод будет вызываться с nil получателем.
В go принято обрабатывать:
type I interface {
	M()
}

type T struct {
	S string
}

"Ошибка времени выполнения, поскольку внутри интерфейса нет типа, указывающего, какой конкретный метод вызывать"
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}


Реализация нескольких интерфейсов!!!!!
При этом тип данных необязательно должен реализовать только методы интерфейса,
для типа данных можно определить его собственные методы или также реализовать методы других интерфейсов.
Пример:

type Reader interface{
    read()
}

type Writer interface{
    write(string)
}

func writeToStream(writer Writer, text string){
    writer.write(text)
}
func readFromStream(reader Reader){
    reader.read()
}

type File struct{
    text string
}

func (f *File) read(){
    fmt.Println(f.text)
}

func (f *File) write(message string){
    f.text = message
    fmt.Println("Запись в файл строки", message)
}

func main() {
    myFile := &File{}
    writeToStream(myFile, "hello world")
    readFromStream(myFile)
}
*/

type Abser interface {
	Abs() float64
}

func printTest(ab Abser) {
	fmt.Println(ab.Abs())
}

type MyFloat float64

// Метод означает, что тип MyFloat реализует интерфейс Abser (не явно)
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyTest struct {
	Test float64
}

func main() {
	println("==========Одиночные интерфейсы========")
	//Объект интерфейса на прямую создать нельзя
	//var ab Abser = Abser{}

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	var _ Abser
	//Нет ошибки
	_ = f
	//Нет ошибки
	_ = &v
	//Ошибка коде так как Vertex не реализуется Abser, так ка метод только для указателя Vertex
	//_ = v

	//Вызов метода, где аргумент интерфейс, а структуры соответствуют интерфейсу
	printTest(f)
	printTest(&v)
	//Вызов метода, где аргумент интерфейс, а структура НЕ соответствует интерфейсу
	//t := MyTest{10}
	//printTest(t)

	println("==========Проверка/утверждение типа========")
	typeAssertions()

	println("==========Переключатель типа========")
	typeSwitches(21)
	typeSwitches("hello")
	typeSwitches(true)

	println("=========Вложенные интерфейсы========")
	myFile := &File{}
	writeToStream(myFile, "hello world")
	readFromStream(myFile)
	writeToStream(myFile, "lolly bomb")
	readFromStream(myFile)
}

//================================================

//Проверка/утверждение типа
// т := i.(Т) - значение интерфейса i содержит конкретный тип Т, и присваивает переменной T базовое значение t
// Если i не содержит T бросится panica
// Данное утверждение может вернуть boolean: т, ок := i.(T)

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64)
	//fmt.Println(f)
}

//================================================

// Переключатель типа
// Тоже самое что и обычный switch, только вместо значения используется type
func typeSwitches(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v = type int\n", v)
	case string:
		fmt.Printf("%v = type string\n", v)
	default:
		fmt.Printf("%v = type %T\n", v, v)
	}
}

//===================================
//Вложенные интерфейсы
//Одни интерфейсы могут содержать другие интерфейсы.
//Для соответствия подобному интерфейсу типы данных должны реализовать все его вложенные интерфейсы

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

type ReaderWriter interface {
	Reader
	Writer
}

func writeToStream(writer ReaderWriter, text string) {
	writer.write(text)
}
func readFromStream(reader ReaderWriter) {
	reader.read()
}

type File struct {
	text string
}

func (f *File) read() {
	fmt.Println(f.text)
}
func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}
