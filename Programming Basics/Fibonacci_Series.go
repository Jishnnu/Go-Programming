package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}


func main() {
	var n int
	fmt.Print("\nENTER A NUMBER FOR GENERATING THE FIBONACCI SERIES : ")
	fmt.Scanln(&n)
	fmt.Println("\nFIBONACCI SERIES :")
	for i := 0; i < n; i++ {
		fmt.Print(fibonacci(i), " ")
	}	
	fmt.Println()
}
