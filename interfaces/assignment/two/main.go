/**
 * @author [mbugua]
 * @create date 2020-06-09 10:56:56
 * @modify date 2020-06-09 10:56:56
 * @desc [description]
 */
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "FILE")
	}
	filename := os.Args[1]
	fmt.Println("file to read:", filename)
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, file)

}
