package main

import (
	"fmt"
	"time"
)

func calcSquare(n int, ans chan int) {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	ans <- sum
}

func calcCube(n int, ans chan int) {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit * digit
		n /= 10
	}
	ans <- sum
}

func main() {
	n := make([]int, 3)
	now := time.Now()
	calcChannel := make(chan int)
	fmt.Print("\nENTER 3 NUMBER : ")
	for i := 0; i < len(n); i++ {
		fmt.Scan(&n[i])
		go calcSquare(n[i], calcChannel)
		time.Sleep(3 * time.Second)
		go calcCube(n[i], calcChannel)
		square, cube := <-calcChannel, <-calcChannel // BUFFERING
		fmt.Println("\nSUM OF SQUARES OF DIGITS OF", n[i], ": ", square)
		fmt.Println("SUM OF CUBES OF DIGITS OF", n[i], ": ", cube)
	}

	close(calcChannel)
	fmt.Println("\nTIME TAKEN : ", time.Since(now))
}
