package main

import "fmt"

func main() {
	var units, price float64
	fmt.Print("\nENTER THE NUMBER OF UNITS : ")
	fmt.Scanln(&units)

	if units <= 50 {
		price = units * 10
	} else if units > 50 && units <= 150 {
		price = 500 + (units-50)*15
	} else if units > 150 && units <= 500 {
		price = 1500 + 500 + (units-150)*25
	} else if units > 500 {
		price = units * 25
	}

	fmt.Println("\nELECTRICITY BILL AMOUNT : RS.", price)
}
