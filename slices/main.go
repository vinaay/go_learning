package main

import "fmt"

func main() {
	var slice []int              // Declare a slice of integers
	slice = []int{1, 2, 3, 4, 5} // Initialize a slice
	slice = append(slice, 6)     // Append to a slice
	fmt.Println(slice[0:3])      // Slice elements from index 0 to 2

}

// to execute us below command
// go run main.go
