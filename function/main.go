package main

import "fmt"

func multiply(a, b string) (result string) {
	result = a + b
	return

}

func main() {
	ans := multiply("ash", "ish")
	fmt.Println("the ans is:", ans)

}
