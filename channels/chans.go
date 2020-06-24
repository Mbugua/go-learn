package main

import "fmt"

func main() {
	c := make(chan string)

	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	// for s := range c {
	// 	fmt.Println(s)
	// }

	for {
		fmt.Println(<-c)
	}
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done Printing."
}
