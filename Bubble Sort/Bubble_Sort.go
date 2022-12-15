// 20191CSE0215

package main

import "fmt"

// FUNCTION THAT APPLIES BUBBLE SORT ON THE INPUT ARRAY
func bubbleSort(arr []int, n int) []int {	
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
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
	fmt.Println("\nSORTED ARRAY : ", bubbleSort(arr, len(arr)))
}
