// 20191CSE0215

package main

import (
	"fmt"
	"os"
)

type BinaryTree struct {
	data  int
	right *BinaryTree
	left  *BinaryTree
}

// FUNCTION THAT INSERTS THE INPUT ELEMENT
func (bt *BinaryTree) Insert(data int) *BinaryTree {
	if bt.data >= data {
		if bt.left == nil {
			bt.left = &BinaryTree{data: data}
			return nil
		}
		return bt.left.Insert(data)
	} else {
		if bt.right == nil {
			bt.right = &BinaryTree{data: data}
			return nil
		}
		return bt.right.Insert(data)
	}
}

// FUNCTION THAT DISPLAYS THE TREE ELEMENTS
func (bt *BinaryTree) Display() {
	if bt == nil {
		return
	}

	bt.left.Display()
	fmt.Print(bt.data, " ")
	bt.right.Display()
}

func main() {
	tree := BinaryTree{}
	ch, i := 0, 1
	for {
		fmt.Print("\n1: INSERT\n2: DISPLAY\n3: EXIT\n\nENTER CHOICE : ")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			fmt.Print("\nENTER AN INTEGER ELEMENT TO INSERT INTO THE BINARY TREE : ")
			fmt.Scan(&i)
			tree.Insert(i)
			break

		case 2:
			fmt.Println("\nBINARY TREE (INORDER TRAVERSAL) : ")
			tree.Display()
			fmt.Println()
			break

		case 3:
			os.Exit(0)

		default:
			fmt.Println("\nINVALID CHOICE")
			break
		}
	}
}
