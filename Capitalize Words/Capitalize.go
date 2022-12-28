package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("\nENTER A SENTENCE WITH ATLEAST FOUR WORDS : ")
	scanner := bufio.NewReader((os.Stdin))
	str, _ := scanner.ReadString('\n')

	count := strings.Count(str, " ") + 1
	if count < 4 {
		fmt.Printf("\nTHE INPUT HAS ONLY %d WORDS. PLEASE ENTER A SENTENCE WITH ATLEAST FOUR WORDS\n", count)
		return
	}

	s := []rune(str)
	for i, v := range s {
		if v >= 'a' && v <= 'z' {
			v -= 32
		}

		if v == ' ' {
			i += 1
			if s[i] >= 'a' && s[i] <= 'z' {
				s[i] -= 32
			}
		} else {
			if v >= 'A' && v <= 'Z' {
				v += 32
			}
		}
	}
	str = string(s)
	fmt.Println("\nCAPITALIZED TEXT : ", str)
}
