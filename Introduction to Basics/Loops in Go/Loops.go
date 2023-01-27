package main

import "fmt"

func main() {
	// TRADITIONAL C STYLE
	fmt.Println("\nLOOPING 10 TIMES : ")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	
	// WHILE STYLE : CONDITION ONLY
	j := 0
	fmt.Println("\nLOOPING 10 TIMES : ")
	for j < 10 {
		fmt.Print(j, " ")
		j++
	}

	// INFINTE LOOP
	
	/*fmt.Println("\nLOOPING 10 TIMES : ")
	for  {
		fmt.Print(j, " ")
	}*/

}
