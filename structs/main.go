package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Smith"}
	var alex person

	alex.firstName = "James"
	alex.lastName = "Bond"
	// fmt.Println(alex)
	fmt.Printf("%+v", alex)

}
