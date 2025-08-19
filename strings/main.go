package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "ashish,singune,mca,"
	parts := strings.Split(data, "i")
	fmt.Println("the parts are :\n", parts)

	data1 := "one two two one three two"
	count := strings.Count((data1), "wo")
	fmt.Println("the count of two is:", count)

	data2 := "    asshish,singune    "
	data2_trim := strings.TrimSpace(data2) //
	fmt.Println("the data2 after removing the space is:", data2_trim)

	data3 := "    asshish      singune    "
	data3_trim := strings.Trim(data3, ",") //
	fmt.Println("the data2 after removing the space is:", data3_trim)
	data4 := "ashish"
	data5 := "singune"
	data6 := strings.Join([]string{data4, data5}, " ")
	fmt.Println("the data after joining is:", data6)

}
