package main

import (
	"fmt"
)

func prime(arr ...int) {
	var primes []int
	flag := false
	for _, v := range arr {
		if v == 0 || v == 1 {
			continue
		}
		for i := 2; i <= v/2; i++ {
			if v%i != 0 {
				flag = true
			} else {
				continue
			}
		}
		if flag == true {
			primes = append(primes, v)
		}
	}
	fmt.Println("\nPRIMES : ", primes)
}

func main() {
	var arr []int
	var i int
	fmt.Println("\nENTER AN ARRAY OF ELEMENTS (ENTER -1 TO STOP) : ")
	for i != -1 {
		fmt.Scan(&i)
		if i != -1 {
			arr = append(arr, i)
		}
	}
	prime(arr...)
}
