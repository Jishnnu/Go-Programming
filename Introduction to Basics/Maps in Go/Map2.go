package main

import "fmt"

func main() {
	patient_record := map[int]int{1001: 21, 1002: 35, 1003: 12, 1004: 64, 1005: 17, 1006: 59, 1007: 43, 1008: 10, 1009: 18, 1010: 72}

	var non_vaccinated = []int{}
	var id, age, ch int
	fmt.Print("\nENTER YOUR ID : ")
	fmt.Scanln(&id)
	fmt.Print("\nARE YOUR VACCINATED ? (0 = NO, 1 = YES): ")
	fmt.Scanln(&ch)
	switch ch {
	case 0:
		fmt.Print("ENTER YOUR AGE : ")
		fmt.Scanln(&age)
		patient_record[id] = age
		non_vaccinated = append(non_vaccinated, id)
		fmt.Println("NON-VACCINATED PATIENT ID : ", non_vaccinated)
		break

	case 1:
		delete(patient_record, id)
		fmt.Println("DELETED PATIENT RECORD ", id)
		fmt.Println("PATIENT RECORD : ", patient_record)
		break

	default:
		fmt.Println("\nINVALID CHOICE")
	}
}
