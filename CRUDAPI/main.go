package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type TODO struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformGetRequest() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", resp.StatusCode)
		return
	}
	fmt.Println("Status Code:", resp.StatusCode)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Data:", string(data))

}

func PerformPostRequest() {

	todo := TODO{
		UserID:    6,
		Title:     "Aashish Singune",
		Completed: true,
	}
	// convert the TODO struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	//convert in string
	jsonstring := string(jsonData)
	fmt.Println("JSON Data:", jsonstring)
	//convert into string data to reader
	jsonreader := strings.NewReader(jsonstring)
	myURL := "https://jsonplaceholder.typicode.com/todos"
	req, err := http.Post("POST", myURL, jsonreader)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer req.Body.Close()
	if req.StatusCode != http.StatusCreated {
		fmt.Printf("Error: received status code %d\n", req.StatusCode)
		return
	}
	fmt.Println("Status Code:", req.StatusCode)
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Data:", string(data))

}
func main() {
	fmt.Println("Performing GET request to fetch TODO item...")
	//PerformGetRequest()
	PerformPostRequest()

}
