package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)                        // basic print statement
	fmt.Printf("Variable type: %T \n", welcome) // print variable type

	reader := bufio.NewReader(os.Stdin) // create a new reader to read input from standard input
	fmt.Print("enter you full name: ")  // prompt user for input
	name, _ := reader.ReadString('\n')  // read input until newline character
	fmt.Println("Your name is:", name)  // print the name entered by user

	fmt.Print("enter your age: ")
	age, _ := reader.ReadString('\n')
	fmt.Println("your age is:", age)

	fmt.Print("enter about yourself: ")
	about, _ := reader.ReadString('\n')
	fmt.Println("about you:", about)

}
