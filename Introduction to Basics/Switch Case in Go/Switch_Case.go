package main

import "fmt"

func main() {
	var char string

	fmt.Print("\nENTER A NUMBER TO GET THE CORRESPONDING DAY (BETWEEN 1 & 7) : ")
	fmt.Scanln(&char)

	switch char {
	case "1":
		fmt.Println("\nMonday")
	case "2":
		fmt.Println("\nTuesday")
	case "3":
		fmt.Println("\nWednesday")
	case "4":
		fmt.Println("\nThursday")
	case "5":
		fmt.Println("\nFriday")
	case "6":
		fmt.Println("\nSaturday")
	case "7":
		fmt.Println("\nSunday")
	default:
		fmt.Println("\nInvalid Input")
	}
}
