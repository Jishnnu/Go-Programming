package main

import (
	"customPackage_Calculator"
	"fmt"
	"os"
)

func main() {
	var choice int
	var a, b float64
	for {
		fmt.Print("\n1: ADDITION\n2: SUBTRACTION\n3: MULTIPLICATION\n4: DIVISION\n5: REMAINDER\n6: FACTORIAL\n7: EXIT : \n\nENTER CHOICE (1-7) : ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("\nENTER TWO NUMBERS : ")
			fmt.Scanln(&a, &b)
			fmt.Println("SUM : ", customPackage_Calculator.Sum(a, b))
			break

		case 2:
			fmt.Println("\nENTER TWO NUMBERS : ")
			fmt.Scanln(&a, &b)
			fmt.Println("DIFFERENCE : ", customPackage_Calculator.Difference(a, b))
			break

		case 3:
			fmt.Println("\nENTER TWO NUMBERS : ")
			fmt.Scanln(&a, &b)
			fmt.Println("PRODUCT : ", customPackage_Calculator.Product(a, b))
			break

		case 4:
			fmt.Println("\nENTER TWO NUMBERS : ")
			fmt.Scanln(&a, &b)
			fmt.Println("QUOTIENT : ", customPackage_Calculator.Quotient(a, b))
			break

		case 5:
			fmt.Println("\nENTER TWO NUMBERS : ")
			fmt.Scanln(&a, &b)
			fmt.Println("REMAINDER : ", customPackage_Calculator.Remainder(a, b))
			break

		case 6:
			var n int
			fmt.Print("\nENTER A NUMBER : ")
			fmt.Scanln(&n)
			fmt.Println("FACTORIAL : ", customPackage_Calculator.Factorial(n))
			break

		case 7:
			os.Exit(0)

		default:
			fmt.Println("\nINVALID OPTION")
		}
	}
}
