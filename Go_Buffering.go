package main

import (
	"fmt"
	"time"
)

func write(channel chan int) {
	for i := 0; i < 5; i++ {
		channel <- i
		fmt.Println("SUCCESSFULLY WROTE", i, "TO CHANNEL")
	}
}

func main() {
	now := time.Now()

	ch := make(chan int, 2)
	go write(ch)

	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("READING VALUE", v)
	}
	
	fmt.Println("\nTIME TAKEN : ", time.Since(now))
}
