package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = [5]string{"cse", "cse4", "cse4@gmail.com", "presidency@outlook.com", "pu@yahoo.com"}
	fmt.Println("\nEMAILS USING STRING FUNCTION : ")
	for _, v := range str {
		if strings.Contains(v, "@") && strings.Contains(v, ".com") {
			fmt.Println(v)
		}
	}

	fmt.Println("\nEMAILS WITHOUT USING STRING FUNCTION : ")
	for i := 0; i < len(str); i++ {
		rune1 := []rune(str[i])
		for _, v := range rune1 {
			if v == '@' {
				fmt.Println(str[i])
			}
		}
	}
}