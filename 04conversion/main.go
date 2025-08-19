// // package main

// // import (
// // 	"bufio"
// // 	"os"
// // )

// // func main() {

// // 	reader := bufio.NewReader(os.Stdin)
// // 	input, _ := reader.ReadString('\n')
// // 	_ = input // Use the input or handle it as needed

// // }

// package main

// import "fmt"

// func main() {

// 	var input string = "ashish"
// 	fmt.Println(input)

// 	var tt = "ashishhhh"
// 	fmt.Println(tt)

// 	var as = 8.22
// 	fmt.Println(as)

// 	const pi = 3.14
// 	fmt.Println(pi) // we can not change the value of pi variable

// 	ashish := 11
// 	fmt.Printf("variable type name :%T\n", ashish)
// 	ashi := "ashish"

// }

package main

import "fmt"

func main() {

	age := 25
	name := "Aashish"
	isStudent := true
	height := 5.999
	pi := 3.14
	complexNum := 1 + 2i

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Name: %s\n, isStudent: %t\n, Height: %.3f\n, pi: %.2f\n, complexNum: %v\n", name, isStudent, height, pi, complexNum)

}
