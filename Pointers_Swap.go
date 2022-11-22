package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	var a, b int
	fmt.Println("\nENTER TWO NUMBERS : ")
	fmt.Scanln(&a, &b)
	swap(&a, &b)
	fmt.Println("\nSWAPPED NUMBERS : ", a, " ", b)
}
