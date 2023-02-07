package main

import (
	"fmt"
	"time"
)

func palindrome(str string) bool {
	flag := 0
	for i := 0; i < len(str)-1; i++ {
		for j := len(str) - 1; j > 0; j-- {
			if str[i] == str[j] {
				flag = 1
			} else {
				flag = 0
			}
		}
	}
	if flag == 1 {
		return true
	}
	return false
}

func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func main() {

	var str string
	var n int
	fmt.Print("\nENTER A STRING : ")
	fmt.Scan(&str)
	fmt.Print("\nENTER AN INTEGER COMPUTE THE FACTORIAL : ")
	fmt.Scan(&n)

	now := time.Now()
	fmt.Println("PALINDROME STATUS : ", palindrome(str))
	fmt.Println("FACTORIAL : ", factorial(n))
	fmt.Println("\nTIME TAKEN FOR EXECUTION : ", time.Since(now))
}
