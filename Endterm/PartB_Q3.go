// 20191CSE0215

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

func fibonacci(ch chan int) {
	a, b := 0, 1
	for {
		ch <- a
		a, b = b, a+b
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	now := time.Now()
	var str string
	var n int
	ch := make(chan int)

	fmt.Print("\nENTER A STRING TO CHECK IF IT IS A PALINDROME : ")
	fmt.Scan(&str)
	fmt.Print("\nENTER AN INTEGER TO COMPUTE THE FIBONACCI SERIES : ")
	fmt.Scan(&n)

	go palindrome(str)
	go fibonacci(ch)
	for i := 0; i <n; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("\nTIME TAKEN FOR EXECUTION : ", time.Since(now))
}
