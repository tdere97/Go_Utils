package main

import "fmt"

func returnfunc() [4]string {
	return [4]string{"Hi", "In", "New", "Fucntion"}

}

func arr() {
	fmt.Println("===============================array.go=================================")
	// Arrays are fixed-size groups of variables of the same type.
	// The type [n]T is an array of n values of type T
	// To declare an array of 10 integers:
	var myInts [10]int
	for i := 0; i < len(myInts); i++ {
		fmt.Printf("%v\t", myInts[i])
	}
	fmt.Printf("\n")

	//or to declare an initialized literal:
	hellos := [4]string{"Hi", "Tejal", "Nivrutti", "Dere"}
	for i := 0; i < len(hellos); i++ {
		fmt.Printf("%v\t", hellos[i])
	}
	fmt.Printf("\n")

	str := returnfunc()
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v\t", str[i])
	}
	fmt.Printf("\n")
	slices()
}
