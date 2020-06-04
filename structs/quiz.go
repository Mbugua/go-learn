package main

import "fmt"

func main() {
	name := "Bill"
	fmt.Println(*&name)
}
