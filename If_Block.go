package main

import "fmt"

func main() {
	var n1 int
	fmt.Print("\nENTER A NUMBER : ")
	fmt.Scanln(&n1)
	if n1 == 0 {
		fmt.Printf("\n%d IS ZERO\n", n1)
	} else if n1 > 0 {
		fmt.Printf("\n%d IS POSITIVE\n", n1)
	} else {
		fmt.Printf("\n%d IS NEGATIVE\n", n1)
	}
}
