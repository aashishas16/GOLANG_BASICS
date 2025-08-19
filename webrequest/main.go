package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("test  a web request")

	responsee, err := http.Get("https://jsonplaceholder.typicode.com/todos/1") // this is a web request to get the data from the url
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	defer responsee.Body.Close() // defer is used to close the response body after the function completes

	// read the response body

	data, err := ioutil.ReadAll(responsee.Body) // io.ReadAll is used to read the response body
	if err != nil {
		fmt.Println("error reading response body:", err)
		return
	}
	fmt.Println("Response data:", string(data))                                   // print the response data as a string
	fmt.Println("Status Code:", responsee.StatusCode)                             // print the status code of the response
	fmt.Println("Status:", responsee.Status)                                      // print the status of the response
	fmt.Println("Headers:", responsee.Header)                                     // print the headers of the response
	fmt.Println("Content Length:", responsee.ContentLength)                       // print the content length of the response
	fmt.Println("Protocol:", responsee.Proto)                                     // print the protocol of the response
	fmt.Println("Request Method:", responsee.Request.Method)                      // print the request method of the response
	fmt.Println("Request URL:", responsee.Request.URL)                            // print the request URL of the response
	fmt.Println("Request Header:", responsee.Request.Header)                      // print the request header of the response
	fmt.Println("Request Host:", responsee.Request.Host)                          // print the request host of the response
	fmt.Println("Request Proto:", responsee.Request.Proto)                        // print the request protocol of the response
	fmt.Println("Request Content Length:", responsee.Request.ContentLength)       // print the request content length of the response
	fmt.Println("Request Transfer Encoding:", responsee.Request.TransferEncoding) // print the request transfer encoding of the response

}
