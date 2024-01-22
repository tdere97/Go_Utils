package main

import "fmt"

func function(a int, b int) int {
	fmt.Println("===============================function.go=================================")
	// func funcname (argument...)(return type){}
	return a + b
}

// Naming return value
func function1(a int, b int) (sum int) {
	// func funcname (argument...)(return type){}
	sum = a + b
	return
}
func nextProgram() {
	structure()
}

//Callback function
// func test (func callBackFuntionName(int,int)(int),int)
// function test takes 2 argument 1st is fucntion( which takes 2 argument both as int and returns int) n second is integer
