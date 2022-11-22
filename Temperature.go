package main

import "fmt"

func main() {
	var temp_f, temp_c float64
	fmt.Print("\nENTER TEMPERATURE IN DEGREE FARANHEIT : ")
	fmt.Scanf("%f", &temp_f)

	temp_c = ((temp_f - 32) )  * 5 / 9
	fmt.Println("TEMPERATURE : ", temp_c, " DEGREE CELCIUS")
}