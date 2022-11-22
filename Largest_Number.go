package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("\nEVEN THREE NUMBERS : ")
	fmt.Scanln(&a, &b, &c)

	if a > b {
		if a > c {
			fmt.Println("\n", a, " IS THE LARGEST NUMBER")
		} else {
			fmt.Println("\n", c, " IS THE LARGEST NUMBER")
		}
	} else if b > a {
		if b > c {
			fmt.Println("\n", b, " IS THE LARGEST NUMBER")
		} else {
			fmt.Println("\n", c, " IS THE LARGEST NUMBER")
		}
	} else {
		fmt.Println("\nUNKNOWN ERROR")
	}
}
