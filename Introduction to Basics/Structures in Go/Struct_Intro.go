package main

import "fmt"

type bank_account struct {
	acc_num int
	name    string
	balance float64
	deposit float64
}

func main() {
	var b1, b2 bank_account

	fmt.Println("\nENTER THE DETAILS OF PERSON 1 :")
	fmt.Print("ACCOUNT NUMBER : ")
	fmt.Scanln(&b1.acc_num)
	fmt.Print("ACCOUNT HOLDER NAME : ")
	fmt.Scanln(&b1.name)
	fmt.Print("ACCOUNT BALANCE : ")
	fmt.Scanln(&b1.balance)
	fmt.Print("DEPOSIT AMOUNT : ")
	fmt.Scanln(&b1.deposit)

	b1.balance += b1.deposit

	fmt.Print("ACCOUNT NUMBER : ")
	fmt.Scanln(&b2.acc_num)
	fmt.Print("ACCOUNT HOLDER NAME : ")
	fmt.Scanln(&b2.name)
	fmt.Print("ACCOUNT BALANCE : ")
	fmt.Scanln(&b2.balance)
	fmt.Print("DEPOSIT AMOUNT : ")
	fmt.Scanln(&b2.deposit)

	b2.balance += b2.deposit

	fmt.Println("\n\nDETAILS OF PERSON 1 :")
	fmt.Printf("ACCOUNT NUMBER : %d\nACCOUNT HOLDER NAME :%s\nBALANCE : RS.%.2f\nDEPOSIT : RS.%.2f\n", b1.acc_num, b1.name, b1.balance, b1.deposit)
	fmt.Println("\nDETAILS OF PERSON 2 :")
	fmt.Printf("ACCOUNT NUMBER : %d\n\nACCOUNT HOLDER NAME :%s\nBALANCE : RS.%.2f\nDEPOSIT : RS.%.2f\n", b2.acc_num, b2.name, b2.balance, b2.deposit)
}
