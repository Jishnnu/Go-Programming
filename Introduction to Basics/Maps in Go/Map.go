package main

import (
	"fmt"
)

func main() {
	max, sum := 0, 0
	var day string
	temp := map[string]int{"Sunday": 32, "Monday": 30, "Tuesday": 29, "Wednesday": 25, "Thursday": 25, "Friday": 27, "Saturday": 24}

	for key, value := range temp {
		if value > max {
			max = value
			day = key
		}
		sum += value
	}

	fmt.Printf("\nHIIGHEST TEMPERATURE : %d\nDAY : %s\n", max, day)
	fmt.Println("AVERAGE TEMPERATURE : ", (sum / 7))

	/*value_1, flag_1 := temp["Wednesday"]
	fmt.Println("IS WEDNESDAY A VALID KEY ? : ", flag_1, value_1)

	value_2, flag_2 := temp["June"]
	fmt.Println("IS JUNE A VALID KEY ? : ", flag_2, value_2)
	strings.EqualFold(day, "Yes")*/
}
