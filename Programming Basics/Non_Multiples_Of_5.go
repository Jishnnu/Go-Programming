package main

import "fmt"

func main() {
	var i int = 0
	fmt.Println("\nNUMBERS BETWEEN 1 & 20 WHICH AREN'T MULTIPLES OF 5 : ")
	for i <= 20 {
		i++
		if i%5 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
}
