package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Bad recommendation of syntax
	//alex := person{"Alex", "Anderson"}

	// One way to assign struct
	//alex := person{firstName: "Alex", lastName: "Anderson"}

	// Another way to assign struct
	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// Not recommended to write this way
	//jimPointer := &jim
	//jimPointer.updateName("jimmy")
	jim.updateName("jimmy")

	jim.print()

}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
