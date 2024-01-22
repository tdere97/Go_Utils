package main

import "fmt"

func slices() {
	fmt.Println("===============================slices.go=================================")
	fmt.Printf("Index of array :\t")
	for i := 0; i < 6; i++ {
		fmt.Printf("%v\t", i)
	}
	fmt.Printf("\n")

	a := [6]int{6, 12, 5, 9, 1, 7}
	fmt.Printf("Initial array :\t\t")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v\t", a[i])
	}
	fmt.Printf("\n")

	// A slice is a dynamically-sized, flexible view of the elements of an array.
	// Slices always have an underlying array, though it isn't always specified explicitly.

	// first index is including and last index is excluding
	// that is 1 is included n 4th is exculeded
	// 2(1st index) ,3(2nd index) ,4(3rd index)
	b := a[1:4]
	fmt.Printf("slice 1 -> 4 :\t")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%v\t", b[i])
	}
	fmt.Printf("\n")

	b = a[0:6]
	fmt.Printf("slice 0 -> 5 :\t")
	for i := 0; i < len(b); i++ {
		fmt.Printf("%v\t", b[i])
	}
	fmt.Printf("\n")

	b[3] = 122
	fmt.Printf("end array :\t\t")
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v\t", a[i])
	}
	fmt.Printf("\n")
	// The syntax is:
	// ->		arrayname[lowIndex:highIndex]
	// ->		arrayname[lowIndex:]
	// ->		arrayname[:highIndex]
	// ->		arrayname[:]

	makefunc()

}
