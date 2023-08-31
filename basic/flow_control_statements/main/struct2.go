package main

import "fmt"

//Пример встраивание отдельным полем:

type contact2 struct {
	email string
	phone string
}

type person2 struct {
	name        string
	age         int
	contactInfo contact2
}

func main() {

	var tom = person2{
		name: "Tom",
		age:  24,
		contactInfo: contact2{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	tom.contactInfo.email = "supertom@gmail.com"

	fmt.Println(tom.contactInfo.email)
	fmt.Println(tom.contactInfo.phone)
}
