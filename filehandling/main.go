package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("error creating file:", err)
		return
	}
	defer file.Close()

	data := " hello, my self aashish"
	_, errors := io.WriteString(file, data)
	if errors != nil {
		fmt.Println("error writing to file:", errors)
		return
	}
	fmt.Println("data written to file successfully")

	file, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	fmt.Println("File content:\n", string(file))

	// Reading the file easy method
	file, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}
	fmt.Println("data read from file:", string(file))
}
