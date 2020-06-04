package main

import "fmt"

func main() {
	// var colors map[string]string

	// declare a map of type string with values of type string
	// colors := map[string]string{
	// 	"red":   "#f0000",
	// 	"green": "#4bf745",
	// }
	// colors := make(map[string]string, 2)
	// colors["white"] = "#fffff"
	// delete(colors, "white")
	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#f0000",
		"green": "#4bf745",
		"white": "#ffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
