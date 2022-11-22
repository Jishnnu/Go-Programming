package main

import "fmt"

func main() {
	var ch string
	fmt.Print("\nENTER A CHARACTER (A / B / C) : ")
	fmt.Scanln(&ch)
	if ch == "A" || ch == "a" {
		fmt.Println("\nAPPLE")
	} else if ch == "B" || ch == "b" {
		fmt.Println("\nBANANA")
	} else if ch == "C" || ch == "c" {
		fmt.Println("\nCOCONUT")
	} else {
		fmt.Println("\nINVALID INPUT")
	}
}
