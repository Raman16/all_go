package main

import "fmt"

type contactInfo struct { //custom type
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	//contact contactInfo
	contactInfo ///equal to contactInfo:contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		}, //Important Note:Here comma is important.Whenever we are declaring multiline structs like this.every single line
		//must have comma,even if it is last property
	}
	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")
	jim.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) { //*person pointer to type person
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)

}
