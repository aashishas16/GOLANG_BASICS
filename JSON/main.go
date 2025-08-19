package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	user := User{Name: "Aashish", Age: 24, Email: "aashish@example.com"} //marshaling

	// Convert struct to JSON
	jsonData, err := json.Marshal(user) //
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData)) //
	// Unmarshal JSON back to struct
	var userFromJSON User
	err = json.Unmarshal(jsonData, &userFromJSON) //
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Name: %s, Age: %d, Email: %s\n", userFromJSON.Name, userFromJSON.Age, userFromJSON.Email) // Print the struct fields
	fmt.Printf("Name: %s, Age: %d, Email: %s\n", userFromJSON.Name, userFromJSON.Age, userFromJSON.Email) // Print the struct fields
}
