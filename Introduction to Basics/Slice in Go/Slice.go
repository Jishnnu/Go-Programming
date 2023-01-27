package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3}
	s1 = append(s1, 4, 5)
	fmt.Println("\nSLICE : ", s1)
	fmt.Println("SLICE LENGTH : ", len(s1))
	fmt.Println("SLICE CAPACITY : ", cap(s1))

	s2 := make([]int, 3, 6)
	fmt.Printf("\nENTER %d VALUES : \n", len(s2))
	for i := range s2 {
		fmt.Scanln(&s2[i], &s1[i])
	}
	fmt.Println("SLICE : ", s2)
	fmt.Println("SLICE LENGTH : ", len(s2))
	fmt.Println("SLICE CAPACITY : ", cap(s2))

	var arr = [3]int{1, 2, 3}
	var s3 = arr[:]
	fmt.Println("\nSLICE : ", s3)
	fmt.Println("SLICE LENGTH : ", len(s3))
	fmt.Println("SLICE CAPACITY : ", cap(s3))
}
