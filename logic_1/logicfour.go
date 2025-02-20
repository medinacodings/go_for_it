package main

import "fmt"

func logicnumberfour() {
	// Define the size of the slice
	n := 10

	// Create a slice of size n
	slice := make([]int, n)

	// Calculate the starting number
	start := n*2 - 1

	// Populate the slice with decreasing odd number
	for i := 0; i < n; i++ {
		slice[i] = start
		start -= 2
	}

	// Print the result
	fmt.Println(slice)
}

func main() {
	// Call logicnumbefour() from main()
	logicnumberfour()
}
