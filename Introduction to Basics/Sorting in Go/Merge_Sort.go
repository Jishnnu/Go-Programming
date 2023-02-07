package main

import "fmt"

func merge(arr []int, low int, mid int, high int) []int {
	n1 := mid - low + 1
	n2 := high - mid

	var lower = make([]int, n1)
	var higher = make([]int, n2)

	for i := 0; i < n1; i++ {
		lower[i] = arr[low+i]
	}

	for i := 0; i < n2; i++ {
		higher[i] = arr[mid+1+i]
	}

	i, j, k := 0, 0, low
	for i < n1 && j < n2 {
		if lower[i] <= higher[j] {
			arr[k] = lower[i]
			k++
			i++
		} else {
			arr[k] = higher[j]
			k++
			j++
		}
	}

	for i < n1 {
		arr[k] = lower[i]
		k++
		i++
	}

	for j < n2 {
		arr[k] = higher[j]
		k++
		j++
	}
	return arr
}

func mergeSort(arr []int, low int, high int) []int {
	if len(arr) <= 1{
		return arr
	}
	if low < high {
		mid := (low + (high - 1)) / 2
		mergeSort(arr, low, mid)
		mergeSort(arr, mid+1, high)
		return merge(arr, low, mid, high)
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
	fmt.Println("\nSORTED ARRAY : ", mergeSort(arr, 0, len(arr)))
}
