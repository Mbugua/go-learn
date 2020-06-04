package main

import "fmt"

func main() {
	// create a slice of integer 0-10
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// sort through the slice and check for even odd numbers
	// even numbers are divisible by 2
	// odd numbers are not divisible by 2
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println("even", number)
		} else {
			fmt.Println("odd", number)
		}
	}

}
