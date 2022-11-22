package main

import "fmt"

func square(n *int) {
	*n *= *n 
	
}

func main() {
	var n int
	fmt.Print("ENTER A NUMBER : ")
	fmt.Scanln(&n)
	square(&n)
	fmt.Printf("\nSQUARE : %d\n", n)
}