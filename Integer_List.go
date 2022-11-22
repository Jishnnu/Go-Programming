package main

import "fmt"

func main() {
	src := make([]int, 10, 12)
	var dest = make([]int, 10)
	fmt.Printf("\nENTER %d VALUES : \n", len(src))
	for i := range src {
		fmt.Scanln(&src[i])
		if src[i] >= 1 && src[i] <= 100 {
			dest[i] = src[i]
		}
	}
	fmt.Println("\nOLD SLICE : ", src)
	fmt.Println("LENGTH : ", len(src))
	fmt.Println("CAPACITY : ", cap(src))
	
	fmt.Println("\nNEW SLICE : ", dest)
	fmt.Println("LENGTH : ", len(dest))
	fmt.Println("CAPACITY : ", cap(dest))
}
