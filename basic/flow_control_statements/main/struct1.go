package main

import "fmt"

//Пример встраивание без отдельного поля:

type contact struct {
	email string
	phone string
}

type person struct {
	name string
	age  int
	contact
}

func main() {

	var tom = person{
		name: "Tom",
		age:  24,
		contact: contact{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	tom.email = "supertom@gmail.com"

	fmt.Println(tom.email)
	fmt.Println(tom.phone)
}
