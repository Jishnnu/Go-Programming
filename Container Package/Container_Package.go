package main

import (
	"container/list"
	"fmt"
)

func main() {
	front, back := 0, 0
	fmt.Print("\nENTER AN ELEMENT TO ADD TO THE FRONT OF THE LIST : ")
	fmt.Scanln(&front)
	fmt.Print("\nENTER AN ELEMENT TO REMOVE FROM THE FRONT OF THE LIST : ")
	fmt.Scanln(&back)

	list := list.New()
	list.PushFront(front)
	list.PushBack(back)
	fmt.Println("\nLIST : ")
	list.InsertAfter(front, list.InsertAfter(front, list.InsertAfter(back, list.Front())))
	for l := list.Front(); l != nil; l = l.Next(){
		fmt.Print(l.Value, " ")
	} 
}
