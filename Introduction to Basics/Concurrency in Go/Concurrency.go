package main

import (
	"fmt"
	"time"
)

func palindrome(str string) {
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
	time.Sleep(10 * time.Millisecond)
	if flag == 1 {
		fmt.Println("PALINDROME STATUS : TRUE")
		return
	}
	fmt.Println("PALINDROME STATUS : FALSE")
}

func factorial(n int, ans chan int) {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}	
	time.Sleep(10 * time.Millisecond)
	ans <- fact
}

func main() {
	now := time.Now()
	var str string
	var n int
	ans := make(chan int)

	fmt.Print("\nENTER A STRING : ")
	fmt.Scan(&str)
	fmt.Print("\nENTER AN INTEGER COMPUTE THE FACTORIAL : ")
	fmt.Scan(&n)

	go palindrome(str)
	go factorial(n, ans)

	temp := <-ans
	fmt.Println("FACTORIAL : ", temp)
	fmt.Println("\nTIME TAKEN FOR EXECUTION : ", time.Since(now))
}
