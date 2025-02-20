package main

import "fmt"

func logicnumbertwelve(n int) {
	// Define the repeating pattern
	pattern := []int{1, 3, 5, 7}

	// Create a slice to hold the sequence
	sequence := make([]int, 0, n)

	// Generate the sequence by repeating the pattern
	for i := 0; i < n; i++ {
		sequence = append(sequence, pattern[i%len(pattern)])
	}

	// Print the sequence
	fmt.Println(sequence)
}

func main() {
	// Call the function n = 12
	fmt.Println("n = 12")
	logicnumbertwelve(12)
}
