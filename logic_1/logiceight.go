package main

import "fmt"

func logicnumbereight(n int) {
	// create the slice for the sequence
	sequence := make([]int, 0, n)

	// Generate the increasing part of the sequence
	for i := 1; i <= (n+1)/2; i++ {
		sequence = append(sequence, i*2)
	}

	// Generate the decreasing part by mirroring the increasing part
	for i := n / 2; i > 0; i-- {
		sequence = append(sequence, i*2)
	}

	// Print the result
	fmt.Println(sequence)
}

func main() {
	// Call logicnumbereight for n=10 and n=11
	fmt.Println("n = 10:")
	logicnumbereight(10)
	fmt.Println("n = 11")
	logicnumbereight(11)
}
