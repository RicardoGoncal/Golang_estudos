package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Ricardo", lastName: "Alves"}

	// fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 06341230,
		},
	}

	// fmt.Println(jim)

	// var alex person
	// alex.firstName = "Ricardo"
	// alex.lastName = "Goncalves"
	// alex.contact.email = "email@gmail.com"
	// alex.contact.zipCode = 06341230

	// fmt.Println(alex)

	// Para exemplo ponteiros
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}
