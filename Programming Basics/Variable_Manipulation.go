package main

import "fmt"

const pi = 3.14
var rad = 10

func main() {
	var name1 string = "Jishnu"
	name2 := " Anil"
	var name3 = " Kumar"
	fmt.Print(name1, name2, name3)
	fmt.Printf("\nName 1 : %T\n", name1) // To get the data type
	fmt.Printf("Name 2 : %T\n", name2)
	fmt.Printf("Name 3 : %T\n", name3)

	var integer int
	var boolean bool
	var floating float64
	var String string
	fmt.Println("\nVARIABLE DECLARATION : ", integer, boolean, floating, String)

	var a, b, c int = 10, 20, 30
	
	d, e, f := 90, 99.99, "Hello World"
	
	var x, y = 100, "Hi"

	fmt.Println("\nALTERNATE METHOD : ", a, b, c, d, e, f, x, y)
	fmt.Println(`\nThis is 
	Go Programming using Back Tick Operator
	instead of New Line Character`)

	var (
		num1 int
		num2 float64 = 9
		str string = "Hello Universe"
	)	

	fmt.Println("\nVARIABLE BLOCK : ",num1, num2, str)
	fmt.Println("\nGLOBAL DECLARATION : ", rad, pi)

	//TAKING USER INPUTS
	var name string
	var age int
	fmt.Print("\nENTER YOUR NAME & AGE : ")
	fmt.Scanln(&name, &age)
	fmt.Println("\nUSER INPUTS : ", name, age)
}