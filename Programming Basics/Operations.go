package main

import ("fmt"; "math")

func main() {
	var num1, num2 float64

	fmt.Print("\nENTER FIRST NUMBER : ")
	fmt.Scanln(&num1)

	fmt.Print("\nENTER SECOND NUMBER : ")
	fmt.Scanln(&num2)

	fmt.Println("\nSUM : ", (num1 + num2))
	fmt.Println("\nMAX ELEMENT : ", math.Max(num1, num2))
	fmt.Println("\nMIN ELEMENT : ", math.Min(num1, num2))	
	fmt.Println("\nSQUARE ROOT OF ", num1, ": ", math.Sqrt(num1))
	fmt.Println("\nCUBE ROOT OF ", num2, ": ", math.Cbrt(num2))
}
