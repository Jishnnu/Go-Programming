package main

import "fmt"

func main() {
	var sum int 
	num := [10] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range num{
		sum += i
	}
	fmt.Println("\nSUM : ", sum)
}