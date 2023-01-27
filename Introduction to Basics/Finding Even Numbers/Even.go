package main

import "fmt"

func main() {
	fmt.Print("\nEVEN NUMBERS BETWEEN 1 & 50 : ")

	for i := 1; i < 50; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
