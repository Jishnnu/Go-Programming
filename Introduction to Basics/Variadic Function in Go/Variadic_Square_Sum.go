package main

import "fmt"

func sum_of_squares(slice ...float64) float64 {
	sum := 0.0
	for _, v := range slice {
		sum += (v * v)
	}
	return sum
}

func main() {
	var slice []float64
	var s float64
	fmt.Println("\nENTER A SET OF FLOATING-POINT NUMBERS (PRESS 0 TO STOP) : ")
	for {
		fmt.Scanln(&s)
		if s == 0 {
			break
		}
		slice = append(slice, s)
	}
	fmt.Println("\nSUM OF SQUARES OF ", slice, " : ", sum_of_squares(slice...))
}
