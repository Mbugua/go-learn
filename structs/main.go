package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
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
		contact: contactInfo{
			email: "jdoe@doe.com",
			phone: "254700123456",
		},
	}

	fmt.Printf("%+v", joe)

}
