package main

import "fmt"

func conditional() {
	fmt.Println("===============================conditional.go=================================")
	x := 5
	// basic if else
	if x > 0 {
		fmt.Printf("Its a Positive Number\n")
	} else {
		fmt.Printf("Its a Negative Number\n")
	}
	// init and condition in one line // the variable scope is limited to the particular if-else block
	if y := 10; y > 0 {
		fmt.Printf("Its a Positive Number in 2nd if condition\n")
	} else {
		y = 1
		fmt.Printf("Made y to : %v\n", y)
	}

	var a, b int
	a = 5
	b = 6
	z := function(a, b)
	fmt.Printf("Addition is function %v\n", z)

	z = function1(a, b)
	fmt.Printf("Addition is function1 %v\n", z)

	nextProgram()

}
