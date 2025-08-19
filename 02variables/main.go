package main

import "fmt"

const Logintoken string = "12122"

func main() {

	var username string = "Aashish"
	fmt.Println("username")
	fmt.Printf("Variable type: %T \n", username)

	var is_true bool = true
	fmt.Println("is_true")
	fmt.Printf("Variable type: %T \n", is_true)

	var check_int uint8 = 255
	fmt.Println("check_int")
	fmt.Printf("Variable type: %T \n", check_int)

	var check_float float32 = 3.14
	fmt.Println("check_float")
	fmt.Printf("variable type: %T \n", check_float)

	var check_complex complex64 = 1 + 2i
	fmt.Println("check_complex")
	fmt.Printf("Variable type: %T \n", check_complex)

	var website = 1.1
	fmt.Println(website)
	fmt.Printf("variable type: %T \n", website)

	number12 := 10
	fmt.Println(number12)

	fmt.Println(Logintoken)
	fmt.Printf("Variable type: %T \n", Logintoken)

}
