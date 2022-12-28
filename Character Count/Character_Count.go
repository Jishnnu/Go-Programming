package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nENTER A SEQUENCE TO PEROFRM WORDCOUNT (TYPE | TO STOP) : ")
	str, _ := reader.ReadString('|')

	dictionary := make(map[string]int)
	for _, v := range str {
		count := 0
		if v != ' ' && v != '|' {
			if dictionary[string(v)] == 0 {
				count++
				dictionary[string(v)] = count
			} else {
				dictionary[string(v)] += 1
			}
		}
	}

	fmt.Println("\nCHARACTER COUNT : ", dictionary)
	fmt.Println("\nCOMPUTE TIME : ", time.Since(now))
}
