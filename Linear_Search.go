// 20191CSE0215

package main

import "fmt"

// FUNCTION THAT SEARCHES FOR THE KEY IN THE INPUT ARRAY
func linearSearch(arr []int, key int) bool {
	for _, v := range arr {
		if v == key {
			return true
		}
	}
	return false
}

func main() {
	var arr []int
	var key int
	i := 1
	fmt.Println("\nENTER AN ARRAY OF INTEGER ELEMENTS TO SORT (PRESS 0 TO STOP) : ")

	// TAKES INPUT AS LONG AS IT IS NOT 0
	for i != 0 {
		fmt.Scan(&i)
		if i != 0 {
			arr = append(arr, i)
		}
	}

	fmt.Print("\nENTER A KEY VALUE TO SEARCH : ")
	fmt.Scan(&key)

	if linearSearch(arr, key) {
		fmt.Println(key, " IS PRESENT IN ", arr)
	} else {
		fmt.Println(key, " IS NOT PRESENT IN ", arr)
	}
}
