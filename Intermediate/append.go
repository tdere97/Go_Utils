package main

import "fmt"

func appendFunc() {
	fmt.Println("===============================append.go=================================")
	// The built-in append function is used to dynamically add elements to a slice:
	// func append(slice []Type, elems ...Type) []Type

	// If the underlying array is not large enough, append() will create a new underlying array and point
	// the slice to it.

	// Notice that append() is variadic, the following are all valid:
	// slice = append(slice, oneThing)
	// slice = append(slice, firstThing, secondThing)
	// slice = append(slice, anotherSlice...)

	sli := make([]int, 5)

	sli[0] = 1
	sli[1] = 3
	sli[2] = 5
	sli[3] = 7
	sli[4] = 9

	fmt.Printf("Slice is :\t")
	for i := 0; i < len(sli); i++ {
		fmt.Printf("%v\t", sli[i])
	}
	fmt.Printf("\n")

	// slice = append(slice, oneThing)
	sli = append(sli, 11)
	fmt.Printf("Slice is :\t")
	for i := 0; i < len(sli); i++ {
		fmt.Printf("%v\t", sli[i])
	}
	fmt.Printf("\n")

	// slice = append(slice, firstThing, secondThing)
	sli = append(sli, 13, 15)
	fmt.Printf("Slice is :\t")
	for i := 0; i < len(sli); i++ {
		fmt.Printf("%v\t", sli[i])
	}
	fmt.Printf("\n")

	// slice = append(slice, anotherSlice...)
	temp := []int{17, 19, 21}
	sli = append(sli, temp...)
	fmt.Printf("Slice is :\t")
	for i := 0; i < len(sli); i++ {
		fmt.Printf("%v\t", sli[i])
	}
	fmt.Printf("\n")

	// dont do this!
	// someSlice = append(otherSlice, element)
	sSlice()
}
