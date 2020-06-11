/**
 * @author [mbugua]
 * @create date 2020-06-09 10:56:56
 * @modify date 2020-06-09 10:56:56
 * @desc [description]
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readFromTextFile(filename string) string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//convert the bytes to string
	txt := string(bs)
	return txt

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "FILE")
	}
	filename := os.Args[1]
	fmt.Println("file to read:", filename)

	tx := readFromTextFile(filename)
	fmt.Println(tx)
}
