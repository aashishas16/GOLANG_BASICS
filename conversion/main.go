package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int = 44
	fmt.Println("the number is:", num)
	fmt.Printf("type of number is: %T\n", num)

	var data float64 = float64(num)
	data = data + 10.55
	fmt.Printf("the data is: %T\n", data)
	fmt.Println("the data is:", data)

	var str string = "1234"
	fmt.Println("the string id:", str)
	fmt.Printf("type of string is: %T\n", str)

	// var data1 int = int(str)
	// fmt.Println("the data1 is :", data1)
	// fmt.Printf("type of data1 is: %T\n", data1)

	data1, _ := strconv.Atoi(str)
	fmt.Println("convert string to int")
	fmt.Println("the data1 is :", data1)
	fmt.Printf("type od data1 is : %T\n", data1)

	var data2 int = 123
	fmt.Println("the data2 is:", data2)
	fmt.Printf("type of data2 is: %T\n", data2)

	var str2 string = strconv.Itoa(data2)
	fmt.Println("convert int to string")
	fmt.Println("the str2 is:", str2)
	fmt.Printf("type of str2 is: %T\n", str2)
}
