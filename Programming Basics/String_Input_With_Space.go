package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Println("\nENTER A STRING : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()
	rune := []rune(s)
	fmt.Printf("\nRUNE REPRESENTATIONS OF %s : \nRUNE LITERAL : %c\nUNICODE : %U\nASCII : %d\n", s, rune, rune, rune)
}
