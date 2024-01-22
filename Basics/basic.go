package main

import "fmt"

func main() {
	fmt.Println("===============================basic.go=================================")
	// single-line comments start with "//"
	// comments are just for documentation - they don't execute
	fmt.Println("hello basics")
	//data types:
	// bool
	// string
	// int int8 int16 int32 int64 -> whole number
	// unit unit8 uint16 uint32 uint64 uintptr
	// byte -> alias of uint8
	// rune -> alias of int32 (represents a unicode)
	// float32 float64 -> decimal number
	// complex64 complex128
	var i int
	var j uint
	var x byte
	var isBool bool
	var s string
	var f float32

	i = -1
	j = 1
	x = 'A'
	isBool = true
	s = "this is just a start"
	f = 32.5

	fmt.Printf("All the variables are i : %v\t j=%v\t x=%v\t isBool=%v\t s=%s\t f=%.1f\n", i, j, x, isBool, s, f)

	//types of variable:
	fmt.Printf("Types of all the variables are i : %T\t j=%T\t x=%T\t isBool=%T\t s=%T\t f=%T\n", i, j, x, isBool, s, f)

	//implicit conversion  and another way to define variable
	temp := int(f)
	fmt.Printf("Implicitly converting float to int temp : %v\n", temp)

	//const variable
	const Pi = 3.14
	const gravity = 3.8
	// if we try to put value in const will pop a error
	// ex. Pi=5.0 -> compiler error

	fmt.Printf(" Value of Pi : %v\t gravity : %v\n", Pi, gravity)

	conditional()
}
