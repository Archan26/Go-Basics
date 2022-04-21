package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Parmar"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Archan"
	// alex.lastName = "Parmar"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Archan",
		lastName:  "Parmar",
		contact: contactInfo{
			email:   "abc@gmail.com",
			zipcode: 334455,
		},
	}
	fmt.Println(jim)
}
