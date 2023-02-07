// 20191CSE0215

package main

import "fmt"

// FUNCTION THAT APPLIES SELECTION SORT ON THE INPUT ARRAY
func selectionSort(arr []int, n int) []int {	
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func main() {
	var arr []int
	i := 1
	fmt.Println("\nENTER AN ARRAY OF INTEGER ELEMENTS TO SORT (PRESS 0 TO STOP) : ")

	// TAKES INPUT AS LONG AS IT IS NOT 0 
	for i != 0 {
		fmt.Scan(&i)
		if i != 0 {
			arr = append(arr, i)
		}
	}
	fmt.Println("\nSORTED ARRAY : ", selectionSort(arr, len(arr)))
}
