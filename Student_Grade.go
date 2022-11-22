package main

import "fmt"

func main() {
	var marks float64
	fmt.Print("\nENTER THE MARKS (OUT OF 100) : ")
	fmt.Scanln(&marks)

	if marks <= 100 && marks > 90 {
		fmt.Println("\nAWARDED GRADE : A+")
	} else if marks <= 90 && marks > 80 {
		fmt.Println("\nAWARDED GRADE : A")
	} else if marks <= 80 && marks > 70 {
		fmt.Println("\nAWARDED GRADE : B+")
	} else if marks <= 70 && marks > 60 {
		fmt.Println("\nAWARDED GRADE : B")
	} else if marks <= 60 && marks > 50 {
		fmt.Println("\nAWARDED GRADE : C+")
	} else if marks <= 50 && marks > 40 {
		fmt.Println("\nAWARDED GRADE : C")
	} else if marks <= 40 && marks > 30 {
		fmt.Println("\nAWARDED GRADE : D")
	} else{
		fmt.Println("\nAWARDED GRADE : F")
	}
}
