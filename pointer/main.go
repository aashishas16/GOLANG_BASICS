package main

import "fmt"

func main() {

	var a int
	a = 10
	var b int
	b = 20

	var ptr *int // pointer to an int
	ptr = &a     // ptr now holds the address of a
	*ptr = 30    // dereferencing ptr to change the value of a
	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Address of a: %p\n", &a) // address of a

	fmt.Printf("Value of ptr: %d\n", b)    // value of b
	fmt.Printf("Address of ptr: %p\n", &b) // address of b
}
