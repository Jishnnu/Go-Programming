package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
}

type Rectangle struct {
	length, breadth float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func (c Circle) area() float64 {
	return math.Pi * c.radius
}

func main() {
	var choice int

	fmt.Print("\n1: RECTANGLE\n2: CIRCLE\n\nENTER YOUR CHOICE : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		var length, breadth float64
		fmt.Print("\nENTER LENGTH (in m) : ")
		fmt.Scanln(&length)
		fmt.Print("ENTER BREADTH (in m) : ")
		fmt.Scanln(&breadth)
		rectangle := Rectangle{length, breadth}
		fmt.Println("AREA : ", rectangle.area(), "sqm")
		break

	case 2:
		var radius float64
		fmt.Print("\nENTER RADIUS (in m) : ")
		fmt.Scanln(&radius)
		circle := Circle{radius}
		fmt.Println("AREA : ", circle.area(), "sqm")
		break

	default:
		fmt.Println("\nINVALID CHOICE")
		break
	}
}
