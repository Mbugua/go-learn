package main

import "fmt"

func main() {
	m := map[string]string{
		"dog": "bark",
		"cat": "purr",
	}
	// changeMap(m)

	// fmt.Println(m)

	for _, value := range m {
		fmt.Println(value)
	}

}

func changeMap(m map[string]string) {
	m["cat"] = "purr"
}
