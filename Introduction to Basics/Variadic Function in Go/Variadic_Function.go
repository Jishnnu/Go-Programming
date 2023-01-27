package main

import "fmt"

func greatest(n ...int) int {
	max := 0
	for _, v := range n {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	slice := []int{100, 20, 30, 50, 40, 90, 60, 80, 70, 10}
	arr := [10]int{100, 20, 30, 50, 40, 90, 60, 80, 70, 10}
	fmt.Println("\nGREATEST ELEMENT IN ", slice, " : ", greatest(slice...))
	fmt.Println("\nGREATEST ELEMENT IN ", arr, " : ", greatest(arr[:]...))
}
