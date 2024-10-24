package main

import "fmt"

func main() {
	var myMap map[string]int     // Declare a map
	myMap = make(map[string]int) // Initialize the map
	myMap["one"] = 1             // Set key-value pairs
	myMap["two"] = 2
	fmt.Println(myMap["one"]) // Access value by key
	delete(myMap, "one")      // Delete a key-value pair
}

// to execute us below command
// go run main.go
