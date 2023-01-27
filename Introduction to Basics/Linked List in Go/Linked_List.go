// 20191CSE0215

package main

import (
	"fmt"
	"os"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// FUNCTION THAT INSERTS THE INPUT ELEMENT
func (ll *LinkedList) Insert(node *Node) {
	if ll.head == nil {
		ll.head = node
		ll.tail = node
		ll.length++
	} else {
		ll.tail.next = node
		ll.tail = node
		ll.length++
	}
}

// FUNCTION THAT DELETES THE INPUT ELEMENT
func (ll *LinkedList) Delete(value int) {
	if ll.head.data == value {
		ll.head = ll.head.next
		ll.length--
		return
	}

	var previous *Node = nil
	current := ll.head

	for current != nil && current.data != value {
		previous = current
		current = current.next
	}

	if current == nil {
		fmt.Println("\nINVALID VALUE : ", value)
		return
	}

	previous.next = current.next
	ll.length--
	fmt.Println("NODE WITH VALUE", value, " DELETED")
}

// FUNCTION THAT DISPLAYS THE LIST ELEMENTS
func (ll *LinkedList) Display() {
	temp := ll.head
	fmt.Println("\nLINKED LIST :")
	for temp != nil {
		fmt.Printf("%v -> ", temp.data)
		temp = temp.next
	}
	fmt.Println("NIL")
}

func main() {
	list := LinkedList{}
	ch, i := 0, 1
	for {
		fmt.Print("\n1: INSERT\n2: DELETE\n3: DISPLAY\n4: EXIT\n\nENTER CHOICE : ")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			fmt.Println("\nENTER INTEGER ELEMENTS TO INSERT INTO THE LINKED LIST (PRESS 0 TO STOP) : ")
			for i != 0 {
				fmt.Scan(&i)
				if i != 0 {
					list.Insert(&Node{data: i})
				}
			}
			break

		case 2:
			fmt.Print("ENTER A VALUE TO DELETE : ")
			fmt.Scan(&i)
			list.Delete(i)
			break

		case 3:
			list.Display()
			break

		case 4:
			os.Exit(0) // DRAWS THE PROGRAM CONTROL OUT OF THE LOOP 

		default:
			fmt.Println("\nINVALID CHOICE")
			break
		}
	}
}
