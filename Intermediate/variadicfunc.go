package main

import "fmt"

func printStrings(names ...string) {
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func print(sli ...int) {
	fmt.Printf("Elements are : ")
	for i := 0; i < len(sli); i++ {
		fmt.Printf("%v\t", sli[i])
	}
	fmt.Printf("\n")
}

func sum(sli ...int) int {
	//inside a func we will treat this varaidic argument as a simple slice
	sum := 0
	print(sli...)
	for i := 0; i < len(sli); i++ {
		sum += sli[i]
	}
	return sum
}
func variadicfunc() {
	fmt.Println("===============================variadicfunc.go=================================")
	// Many functions, especially those in the standard library, can take an arbitrary number of final arguments.
	// This is accomplished by using the "..." syntax in the function signature.
	// A variadic function receives the variadic arguments as a slice.
	t1 := sum(1, 2, 3)
	fmt.Printf("Sum is : %v\n", t1)

	t2 := sum(1, 2, 3, 5, 6)
	fmt.Printf("Sum is : %v\n", t2)

	t3 := sum(1)
	fmt.Printf("Sum is : %v\n", t3)

	//	SPREAD OPERATOR
	//
	// The spread operator allows us to pass a slice into a variadic function.
	// The spread operator consists of three dots following the slice in the function call.

	names := []string{"bob", "sue", "alice"}
	printStrings(names...)

	appendFunc()

}
