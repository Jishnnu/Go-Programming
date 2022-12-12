// 20191CSE0215

package main

import "fmt"

// FUNCTION THAT APPLIES SELECTION SORT ON THE INPUT ARRAY
func sort(arr []int) []int {
	for i := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
	return arr
}

// FUNCTION THAT SEARCHES FOR THE KEY IN THE INPUT ARRAY
func binarySearch(arr []int, key int) bool {
	low := 0
	high := len(arr) - 1
	arr = sort(arr)
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == key {
			return true
		} else if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	var arr []int
	var key int
	i := 1
	fmt.Println("\nENTER AN ARRAY OF INTEGER ELEMENTS TO SORT (PRESS 0 TO STOP) : ")
	for i != 0 {
		fmt.Scan(&i)
		if i != 0 {
			arr = append(arr, i)
		}
	}

	fmt.Print("\nENTER A KEY VALUE TO SEARCH : ")
	fmt.Scan(&key)

	if binarySearch(arr, key) {
		fmt.Println(key, " IS PRESENT IN ", arr)
	} else {
		fmt.Println(key, " IS NOT PRESENT IN ", arr)
	}
}
