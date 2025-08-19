package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	month, _ := reader.ReadString('\n')

	x := month
	switch x {

	case "January", "February", "March":
		fmt.Println("this is a winter month")

	case "April", "May", "June":
		fmt.Println("this is a summer month")
	case "July", "August", "September":
		fmt.Println("this is a monsoon month")
	case "October", "November", "December":
		fmt.Println("this is a autumn month")
	default:
		fmt.Println("this is not a valid month")
		fmt.Println("please enter a valid month")
	}

}
