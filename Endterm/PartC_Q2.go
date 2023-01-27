// 20191CSE0215

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nENTER ANY NAME : ")
	name, _ := reader.ReadString('\n')

	if len(name) == 0 {
		return 
	}

	name = strings.ToUpper(name)
	fmt.Printf("%c.", (rune)(name[0]))
	rune1 := []rune(name)
	for i, v := range rune1 {
		if v == ' ' && v != ',' {
			fmt.Printf("%c.", rune1[i+1])
			continue
		}
	}

}
