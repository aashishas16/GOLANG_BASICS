package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range numbers {
		fmt.Printf("Index:%d,Value:%d\n", index, value)

	}

	data := "ashish,sing"
	for x, y := range data {
		fmt.Printf("index:%d,Value:%c\n", x, y)
	}
}
