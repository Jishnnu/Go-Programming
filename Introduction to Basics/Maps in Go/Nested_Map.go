package main

import "fmt"

func main() {
	class := map[string]map[string]float64{
		"7CSE-01": {"20191CSE0001": 6.8, "20191CSE0002": 7.2, "20191CSE0003": 8.0},
		"7CSE-02": {"20191CSE0010": 7.8, "20191CSE0011": 6.2, "20191CSE0012": 9.0},
		"7CSE-03": {"20191CSE0100": 8.2, "20191CSE0101": 7.7, "20191CSE0103": 6.6},
		"7CSE-04": {"20191CSE0201": 9.3, "20191CSE0202": 8.9, "20191CSE0203": 8.4},
		"7CSE-05": {"20191CSE0301": 8.3, "20191CSE0302": 6.9, "20191CSE0303": 7.4}}

	var roll_no, sect string
	cgpa := 0.0
	fmt.Println("\nENTER ROLL NUMBER : ")
	fmt.Scanln(&roll_no)

	for section := range class {
		for roll_number, score := range class[section] {
			if roll_no == roll_number {
				sect = section
				cgpa = score
			}
		}
	}
	if cgpa != 0.0 {
		fmt.Println("\nSECTION : ", sect, "\nCGPA : ", cgpa)
	} else {
		fmt.Println("\nERROR : INVALID ROLL NUMBER. PLEASE SUBMIT A VALID ENTRY")
	}
}
