package main

import "fmt"

func rangeFunc() {
	fmt.Println("===============================range.go=================================")
	// Go provides syntactic sugar to iterate easily over elements of a slice:
	// for INDEX, ELEMENT := range SLICE {
	// }

	fruits := []string{"apple", "banana", "cherry", "dragonfruit"}

	for i, fruit := range fruits {
		fmt.Printf("At %v index value is : %v\n", i, fruit)
	}
}
