package main

import "fmt"

type Flower interface {
	description() (float64, float64)
}

type Lotus struct {
	petal, leaves float64
}

type Rose struct {
	petal, thorns float64
}

func (l *Lotus) description() (float64, float64) {
	return l.petal, l.leaves
}

func (r *Rose) description() (float64, float64) {
	return r.petal, r.thorns
}

func main() {
	var choice int

	fmt.Print("\n1: LOTUS\n2: ROSE\n\nENTER YOUR CHOICE : ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		var petal, leaf float64
		fmt.Print("\nENTER PETAL LENGTH (in cm) : ")
		fmt.Scanln(&petal)
		fmt.Print("ENTER LEAF LENGTH (in cm) : ")
		fmt.Scanln(&leaf)
		lotus := Lotus{petal, leaf}
		petal, leaf = lotus.description()
		fmt.Println("\nDESCRIPTION :\nPETAL LENGTH : ", petal, " cm\nLEAF LENGTH : ", leaf, " cm")
		break

	case 2:
		var petal, thorn float64
		fmt.Print("\nENTER PETAL LENGTH (in cm) : ")
		fmt.Scanln(&petal)
		fmt.Print("ENTER THORN LENGTH (in cm) : ")
		fmt.Scanln(&thorn)
		rose := Rose{petal, thorn}
		petal, thorn = rose.description()
		fmt.Println("\nDESCRIPTION :\nPETAL LENGTH : ", petal, " cm\nTHORN LENGTH : ", thorn, " cm")
		break

	default:
		fmt.Println("\nINVALID CHOICE")
		break
	}
}
