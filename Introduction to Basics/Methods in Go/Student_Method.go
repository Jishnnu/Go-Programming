package main

import "fmt"

type Student struct {
	name  string
	age   int
	marks [10]float64
}

func (s *Student) averageMarks() float64 {
	avg := 0.0
	for _, v := range s.marks {
		avg += v
	}
	return avg / (float64)(len(s.marks))
}

func (s *Student) getDetails() {
	fmt.Println("STUDENT NAME : ", s.name)
	fmt.Println("STUDENT AGE : ", s.age)
	fmt.Println("STUDENT MARKS : ", s.averageMarks())
}

func main() {
	var name string
	var age int
	var marks [10]float64
	var a float64

	fmt.Print("\nENTER STUDENT NAME : ")
	fmt.Scanln(&name)
	fmt.Print("ENTER STUDENT AGE : ")
	fmt.Scanln(&age)
	fmt.Println("ENTER MARKS OF 10 STUDENTS (ENTER 999 TO STOP) : ")	
	for i := 0; i < 9; i++ {
		fmt.Scanln(&a)
		marks[i] = a
	}
	s := Student{name, age, marks}
	fmt.Println("\nSTUDENT DETAILS : ")
	s.getDetails()
}
