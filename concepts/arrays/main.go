package main

import "fmt"

func main() {
	var arr [5]int              // Declare an array of integers with size 5
	arr = [5]int{1, 2, 3, 4, 5} // Initialize array
	fmt.Println(arr[0])         // Access the first element
}

// Definition: Fixed-size collection of elements of the same type.

// Characteristics:
// 	Fixed size, cannot be resized.
// 	Zero-based indexing.

// to execute us below command
// go run main.go
// test
