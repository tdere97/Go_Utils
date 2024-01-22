package main

import "fmt"

func makefunc() {
	fmt.Println("===============================make.go=================================")
	// func make([]T, len, cap) []T
	//type of array, slice lenght, underlying arrays cpacity
	mySlice1 := make([]int, 5, 10)
	// the capacity argument is usually omitted and defaults to the length
	mySlice2 := make([]int, 5)

	fmt.Printf("Slice 1 :\t")
	for i := 0; i < len(mySlice1); i++ {
		fmt.Printf("%v\t", mySlice1[i])
	}
	fmt.Printf("\n")

	fmt.Printf("Slice 2 :\t")
	for i := 0; i < len(mySlice2); i++ {
		fmt.Printf("%v\t", mySlice2[i])
	}
	fmt.Printf("\n")

	mySlice3 := []string{"Hi!", "Hello!", "Hey!"}

	fmt.Printf("Slice 3 :\t")
	for i := 0; i < len(mySlice3); i++ {
		fmt.Printf("%v\t", mySlice3[i])
	}
	fmt.Printf("\n")
	variadicfunc()
}
