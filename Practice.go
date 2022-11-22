package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c *Circle) Area() float64 {
	return (math.Pi * c.r * c.r)
}

func main() {
	var a float64
	fmt.Print("\nENTER RADIUS : ")
	fmt.Scanln(&a)
	c := Circle{a} // OR var c Circle & fmt.Scanln(&c.r)
	fmt.Printf("\nAREA OF CIRCLE = %.2f\n", c.Area())
}
