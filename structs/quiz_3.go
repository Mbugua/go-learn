package main

import "fmt"

func main() {
	name := "Bill"

	namePointer := &name

	fmt.Println((&namePointer))
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
