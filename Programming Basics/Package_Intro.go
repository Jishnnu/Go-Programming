package main

import (
	"fmt"
	"strings"
	"time"
	"reflect"
)

func main() {
	str1 := "Go"
	str2 := "Programming"
	str3 := "Course"

	//TIME PACKAGE
	fmt.Println("\nCOUNTDOWN : ")
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
		
	}

	fmt.Print("\nSTRING 3 -> INDEX 1 (ASCII) : ", str3[1])
	fmt.Printf("\nSTRING 3 -> INDEX 1 : %c", str3[1])
	fmt.Print("\nSTRING 3 LENGTH : ", len(str3), "\n")

	//STRING FUNCTIONS
	fmt.Println("\nSTRING COMPARISON (STR1 & STR2) : ", strings.Compare(str1, str2))
	fmt.Println("STRING CONTAINS M (STR2) ? : ", strings.Contains(str2, "m"))
	fmt.Println("STRING REPLACE (STR3) ? : ", strings.Replace(str3, "e", "e", 1))
	fmt.Println("STRING EQUALFOLD  (STR1 & Go): ", strings.EqualFold(str1, "Go"))

	//REFLECT PACKAGE
	fmt.Println("TYPE OF STR1 : ", reflect.TypeOf(str1))
}
