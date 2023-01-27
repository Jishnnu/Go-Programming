package main

import "fmt"

func odd_even(n int) (int, bool) {
	val := n / 2
	if n%2 == 0 {
		return val, true
	}
	return val, false
}

func main() {
	var n int
	fmt.Print("\nENTER AN INTEGER : ")
	fmt.Scanln(&n)
	val, flag := odd_even(n)
	fmt.Printf("\nINTEGER HALVED : (%d, %t)\n", val, flag)
}
