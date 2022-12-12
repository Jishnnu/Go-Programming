package main

import "fmt"

func insertionSort(arr []int, n int) []int {
	for i := range arr {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	var arr []int
	i := 1
	fmt.Println("\nENTER AN ARRAY OF INTEGER ELEMENTS TO SORT (PRESS 0 TO STOP) : ")
	for i != 0 {
		fmt.Scan(&i)
		if i != 0 {
			arr = append(arr, i)
		}
	}
	fmt.Println("\nSORTED ARRAY : ", insertionSort(arr, len(arr)))
}
