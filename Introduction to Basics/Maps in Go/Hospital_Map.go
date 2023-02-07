package main

import "fmt"

func main() {
	patient_record := map[int]int{1001: 21, 1002: 35, 1003: 12, 1004: 64, 1005: 17, 1006: 59, 1007: 43, 1008: 10, 1009: 18, 1010: 72}

	var adult_list = []int{}
	var young_list = []int{}

	for key, value := range patient_record {
		if value < 18 {
			young_list = append(young_list, key)
		} else {
			adult_list = append(adult_list, key)
		}
	}

	fmt.Println("\nYOUNG PATIENTS : ", young_list)
	fmt.Println("ADULT PATIENTS : ", adult_list)
}
