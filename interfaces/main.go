package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(sp)
	printGreeting(eb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// custom logic for generating an english greeting
	return "Hi There!"
}

func (sp spanishBot) getGreeting() string {
	return "Hola!"
}
