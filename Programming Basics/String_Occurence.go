package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)
	fmt.Println("\nENTER A STRING (PRESS ENTER TO STOP) : ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	var words []string = strings.Split(str, " ")

	for _, k := range words {
		cnt := 0
		for _, v := range words {
			if k == v {
				cnt++
				count[k] = cnt
			}
		}
	}

	fmt.Println("\nSTRING OCCURENCE : ", count)
}
