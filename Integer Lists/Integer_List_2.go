package main

import (
	"fmt"
)

func main() {
	var arr = []int{}
	a, i := 1, 0
	fmt.Println("\nENTER VALUES FOR THE LIST : ")
	for a != 0 {
		fmt.Scanln(&a)
		arr = append(arr, a)
		i++
	}

	var square = make([]int, len(arr))
	for i, v := range arr {
		square[i] = v * v
	}

	var positive = make([]int, len(arr))
	for i, v := range arr {
		if v >= 0 {
			positive[i] = v
		}
	}

	fmt.Println("\nORIGINAL LIST : ", arr)
	fmt.Println("SQUARED  LIST : ", square)
	fmt.Println("POSITIVE NUMBERED LIST : ", positive)
}
