package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
	// contact contactInfo
}
type contactInfo struct {
	email string
	phone string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Smith"}
	// var alex person

	// alex.firstName = "James"
	// alex.lastName = "Bond"
	// // fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	joe := person{
		firstName: "Joe",
		lastName:  "Doe",
		contactInfo: contactInfo{
			email: "jdoe@doe.com",
			phone: "254700123456",
		},
	}

	joe.updateName("Gi-Joe")
	joe.print()

}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// create person  receiver
func (p person) print() {
	fmt.Printf("%+v", p)
}
