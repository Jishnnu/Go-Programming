package main

import "fmt"

func main() {
	slice := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := 9999
	fmt.Println("\nSLICE : ", slice)
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	fmt.Println("SMALLEST NUMBER : ", min)
}
