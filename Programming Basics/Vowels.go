package main

import "fmt"

func main() {
	var str string
	count := 0
	fmt.Print("\nENTER A STRING : ")
	fmt.Scanf("%s", &str)
	for _, i := range str {
		if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' || i == 'A' || i == 'E' || i == 'I' || i == 'O' || i == 'U' {
			fmt.Printf("%c ", str[i])
			count++
		}
	}
	fmt.Println("\nVOWELS : ", count)
}
