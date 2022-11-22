package main

import (
	"fmt"
)

func main() {
	var s1, s2, sum1, sum2 int
	fmt.Println("\nENTER THE SIZE OF TWO LISTS : ")
	fmt.Scanln(&s1) 
	fmt.Scanln(&s2)

	list1 := make([]int, s1)
	fmt.Printf("\nENTER %d VALUES FOR LIST 1 : \n", len(list1))
	for i := range list1 {
		fmt.Scanln(&list1[i])
	}

	list2 := make([]int, s2)
	fmt.Printf("\nENTER %d VALUES FOR LIST 2 : \n", len(list2))
	for i := range list2 {
		fmt.Scanln(&list2[i])
	}

	if len(list1) == len(list2) {
		fmt.Println("\nBOTH THE LISTS HAVE SAME LENGTHS")
	} else {
		fmt.Println("\nBOTH THE LISTS HAVE DIFFERENT LENGTHS")
	}

	for _, i := range list1 {
		sum1 += i
	}
	for _, i := range list2 {
		sum2 += i
	}

	if sum1 == sum2 {
		fmt.Println("\nBOTH THE LISTS ADD UP TO ", sum1)
	} else {
		fmt.Println("\nBOTH THE LISTS ADD UP TO DIFFERENT VALUES")
	}

	var arr []int
	for _, i := range list1 {
		for _, j := range list2 {
			if i == j {
				arr = append(arr, i)
			}
		}
	}
	if len(arr) > 0{
		fmt.Println("\nSIMILAR ELEMENTS IN BOTH LISTS : ", arr)
	}else{
		fmt.Println("\nTHERE ARE NO SIMILAR ELEMENTS IN THE LISTS")
	}
}
