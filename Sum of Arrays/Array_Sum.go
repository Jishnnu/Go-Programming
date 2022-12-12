package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr [5]int
	j, sum := 0, 0

	//ARRAY INPUT
	fmt.Printf("\nENTER %d ARRAY ELEMENTS : \n", len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Scanln(&j)
		arr[i] = j
	}

	//ARRAY SUMMATION
	for _, v := range arr {
		sum += v
	}

	//PRINTING ARRAY
	fmt.Println("\nTYPE : ", reflect.TypeOf(arr))
	fmt.Println("ARRAY : ", arr)
	fmt.Println("SUM : ", sum)
}
