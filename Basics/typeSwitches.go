package main

import "fmt"

// A type switch is smiliar to a regular switch statement, but the case specfy types instead of values
func printNum(num interface{}) {
	fmt.Println("===============================typeSwitches.go=================================")
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	case float64:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
	forLoop()
}
