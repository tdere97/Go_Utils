package main

import "fmt"

// struct method in go
type rect struct {
	length int
	width  int
}

// struct methods: are just fucntion which have a receiver
// a receiver is a special parameter that syntactically goes before the name of fucntion
func (r rect) area() int {
	return r.length * r.width
}

type mystruct struct {
	name   string
	rollno int
	marks  float32
}

// Nested Structure
type car struct {
	length     int
	width      int
	frontwheel wheel
	backwheel  wheel
}
type wheel struct {
	radius int
}

func test(x mystruct) {
	fmt.Printf("Student details are:\nname : %v\nrollno : %v\nmarks : %v\n", x.name, x.rollno, x.marks)
}
func printingCar(x car) {
	fmt.Printf("Car details are:\nlength : %v\nwidth : %v\nfrontWheel Radius : %v\nbackWheel Radius : %v\n", x.length, x.width, x.frontwheel.radius, x.backwheel.radius)
}

func structure() {
	fmt.Println("===============================struct.go=================================")
	myInstanceofstuct := mystruct{
		name:   "Tejal",
		rollno: 1,
		marks:  90,
	}
	test(myInstanceofstuct)
	//directly assigning instead of creating a instance
	test(mystruct{
		name:   "Dere",
		rollno: 2,
		marks:  91,
	})

	c1 := car{
		length: 12,
		width:  14,
		frontwheel: wheel{
			radius: 5,
		},
		backwheel: wheel{
			radius: 5,
		},
	}
	printingCar(c1)

	//Anonymous Structure in Go
	//It is just like a normal struct , but it can be defined without name and therefore cannot be refferenced elsewhere in the code

	mycar := struct {
		make  string
		model string
	}{
		make:  "Tesla",
		model: "V1",
	}
	fmt.Printf("Car details from Anonymous struct are:\nmake : %v\nmodel : %v\n", mycar.make, mycar.model)

	//we can even nest anonymous struct
	type nestAnonymous struct {
		value     int
		s         string
		anonymous struct {
			v  int
			st string
		}
	}

	a := nestAnonymous{
		value: 1,
		s:     "here we are",
		anonymous: struct {
			v  int
			st string
		}{
			v:  2,
			st: "here we are in anonymous",
		},
	}

	fmt.Printf("Random details from Anonymous struct are:\nvalue : %v\nstring: %v\nvalue : %v\nstring: %v\n", a.value, a.s, a.anonymous.v, a.anonymous.st)

	//Embedded Structs
	//they provide kind of data-only inheritance that can be useful at times
	//Note : go doesnt support classes or inhertance in the complete sense,
	// embedded structs are just a way to elevate and share fields between struct defination

	type vehicle struct {
		make  string
		model string
	}

	type truck struct {
		vehicle
		length int
	}

	t1 := truck{
		vehicle: vehicle{
			make:  "Benz",
			model: "V2",
		},
		length: 40,
	}
	fmt.Printf("Truck details from Embedded struct are:\nmake : %v\nmodel : %v\nlength : %v\n", t1.make, t1.model, t1.length)

	rt := rect{
		length: 10,
		width:  20,
	}
	fmt.Printf("Area of Rectangle is :%v\n", rt.area())

	inter()

}
