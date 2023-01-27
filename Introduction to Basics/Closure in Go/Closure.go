package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("\nENTER TWO NUMBERS : ")
	fmt.Scanln(&a, &b)
	multipy := func(x, y int) int {
		return x * y
	}
	fmt.Println("\nMULTIPLICATION USING CLOSURE: ", multipy(a, b))
}
