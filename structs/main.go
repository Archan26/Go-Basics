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
	//we can write only "contactInfo"
}

func main() {
	//method #1 declaring struct
	// alex := person{"Alex", "Parmar"}
	// fmt.Println(alex)

	//method #2 declaring struct
	// var alex person
	// alex.firstName = "Archan"
	// alex.lastName = "Parmar"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	//Embedded struct
	jim := person{
		firstName: "Archan",
		lastName:  "Parmar",
		contact: contactInfo{
			email:   "abc@gmail.com",
			zipcode: 334455,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Nidhi")
	jim.print()

	//We can write this too
	// jim.updateName("Nidhi")
	// jim.print()
}

func (pointerToPerson *person) updateName(newFirstname string) {
	(*pointerToPerson).firstName = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
