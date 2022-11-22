package main

import (
	"fmt"
	"strings"
)

var fare = map[string]int{"FIRST": 300, "SECOND": 200, "ECONOMY": 100}
var seat = map[string]int{"FIRST": 50, "SECOND": 50, "ECONOMY": 50}
var price, flag, seat_status int
var seat_allotted []string
var seat_class string

func bill(class *string, n *int) {
	for key, val := range fare {
		if strings.EqualFold(*class, key) {
			price += (*n * val)
		}
	}
	for key := range seat {
		if seat[key] == 0 {
			flag++
		}
	}
	invoice()
}

func seat_availability(class *string, people *int) bool {
	for key := range seat {
		if strings.EqualFold(*class, key) {
			val := seat[key]
			if val == 0 {
				fmt.Println("SORRY NO SEATS AVAILABLE IN ", key, " CLASS")
				return false
			} else if *people > val {
				fmt.Println("SORRY ", *people, " SEATS NOT AVAILABLE IN ", key, " CLASS\nPLEASE TRY AGAIN IN SECOND/ECONOMY CLASSES")
				return false
			} else if flag == 3 {
				fmt.Println("\nSORRY HOUSEFULL\nPLEASE TRY AGAIN TOMORROW")
				return false
			}
		}
	}
	return true
}

func seat_allocation(class *string, n *int) {
	people := *n
	if seat_availability(class, &people) {
		for key := range seat {
			val := seat[key]
			if people <= val && (strings.EqualFold(*class, key) || (flag > 0 && flag < 3)) {
				for people > 0 {
					seat_status++
					seat_received := fmt.Sprintf("%s%d", key[0:1], seat_status)
					seat_allotted = append(seat_allotted, seat_received)
					people--
				}
				val -= seat_status
				if val > 0 {
					seat[key] = val
				} else {
					seat[key] = 0
				}
				break
			}			
		}
	}
	bill(class, n)
	fmt.Print(seat, flag)
}

func invoice() {
	fmt.Println("\n****** BOOKING CONFIRMED ******\nSEATS : ", seat_allotted, "\nPRICE : RS.", price)
}

func main() {
	var class string
	var n, cont int
	for {
		fmt.Print("\nENTER THE TICKET CLASS (FIRST/SECOND/ECONOMY) : ")
		fmt.Scanln(&class)

		fmt.Print("\nENTER THE NUMBER OF PEOPLE : ")
		fmt.Scanln(&n)
		seat_allocation(&class, &n)

		fmt.Println("\n\nWANT TO BOOK MORE TICKETS ? (1 == YES, 0 == NO) : ")
		fmt.Scanln(&cont)
		if cont == 0 {
			break
		}
	}
}
