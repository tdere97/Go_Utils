package main

import (
	"fmt"
	"math"
)

// Interfaces in go are collection of method signature
type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.width)
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Function that takes a shape interface and prints its details
func printShapeDetails(sh shape) {
	fmt.Printf("Area: %f\n", sh.area())
	fmt.Printf("Perimeter: %f\n", sh.perimeter())
}
func inter() {
	fmt.Println("===============================interfaces.go=================================")
	rt := rectangle{
		length: 10.3,
		width:  2.4,
	}
	cr := circle{
		radius: 4.1,
	}
	printShapeDetails(rt)
	printShapeDetails(cr)
	printNum(1)
	printNum("tej")
	printNum(1.24)
	printNum(struct{}{})

}

//Interface Implementation
// 1. Interfaces are implemented implicitly
// 2. type (struct) never declares that it implements a given interface
// 3. If an interface exists and a type has the proper methods defined, then the type automatically fullfills that interface
//ex. type circle struct implemets shape{} -> this is done implicitily no need to s=write in this  way

// //Multiple Interfaces
// 1. A type can implement any number of interfaces in GO.
// 2. for ex. the empty interface
// 			type xyz interface{

// 			}
//  always implemented by every type beacuse it has no requirement.

//Name interface argumnets
// Less Readability and clarity
// type copy interface{
// 	cpy(string,string) int
// }
// More Readability and clarity
// type copy interface{
// 	cpy(sourcefile string, destinationfile string) (bytecopied int)
// }

//Type Assertions in Go
// 1. when working with interfaces in GO, every once in a while you will need to access to the underlying type of the interface value.
// 2. you can cast an interface to its undelying type using a type Assertions
// ex. type shape interface{
// 	area()float64
// }
// type circle struct{
// 	radius float64
// }

// c,ok:=s.circle

// -> c is new circle cast from s (which is an instance of shape)
// -> ok is bool that is true if s was a circle or false if s isnt a circle
