package main

import (
	"fmt"
	"net/url"
)

func main() {
	myURL := "https://www.example.com/path/to/resource?query=123&other=456" // Example URL
	fmt.Printf("Original URL:%T", myURL)                                    // Print the type of myURL
	fmt.Println(myURL)                                                      // Print the original URL
	fmt.Println("Handling URL in Go")                                       // Print a message indicating URL handling

	parsedURL, err := url.Parse(myURL) // Parse the URL
	if err != nil {
		fmt.Println("Error parsing URL:", err) // Print error if parsing fails
		return                                 // Exit the program if there is an error
	}
	fmt.Println("Parsed URL:", parsedURL)        // Print the parsed URL
	fmt.Println("Scheme:", parsedURL.Scheme)     // Print the scheme (e.g., "https")
	fmt.Println("Host:", parsedURL.Host)         // Print the host (e.g., "www.example.com")
	fmt.Println("Path:", parsedURL.Path)         // Print the path (e.g., "/path/to/resource")
	fmt.Println("Query:", parsedURL.RawQuery)    // Print the query string (e.g., "query=123&other=456")
	fmt.Println("Fragment:", parsedURL.Fragment) // Print the fragment (if any)
	fmt.Println("User:", parsedURL.User)         // Print the user info (if any)

}
