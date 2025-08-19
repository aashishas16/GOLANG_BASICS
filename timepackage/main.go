package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current time is:", currentTime)

	formatedtime := currentTime.Format("02-01-2006")
	fmt.Println("Formated time is:", formatedtime)
	formatedtime = currentTime.Format("02-01-2006 15:04:05")
	fmt.Println("Formated time is:", formatedtime)

}
