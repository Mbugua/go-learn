package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, error := os.Open(fileName)

	if error != nil {
		fmt.Println(error)
	} else {
		io.Copy(os.Stdout, file)
	}
}
