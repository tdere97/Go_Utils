package main

import "fmt"

// for INITIAL; CONDITION; AFTER{
// 	// do something
//   }

//   ->	INITIAL is run once at the beginning of the loop and can create
//   		variables within the scope of the loop.

//   ->	CONDITION is checked before each iteration. If the condition doesn't pass
//   		then the loop breaks.

//   ->	AFTER is run after each iteration.

func forLoop() {
	fmt.Println("===============================forLoop.go=================================")
	num := 10
	// for i := 1; i <= num; i++ {
	// 	for j := 2; j < i; j++ {
	// 		if i%j != 0 {
	// 			fmt.Printf("%v is prime\n", i)
	// 			break
	// 		} else {
	// 			break
	// 		}
	// 	}
	// }
	for i := 2; i <= num; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%v is prime\n", i)
		}
	}

}

//   ->		the CONDITION (middle part) can be omitted which causes the loop to run forever.
// for INITIAL; ; AFTER {
//   // do something forever
// }

//   ->		a while loop is just a for loop that only has a CONDITION.
// for CONDITION {
//   // do some stuff while CONDITION is true
// }
