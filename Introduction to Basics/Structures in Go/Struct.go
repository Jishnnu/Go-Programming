package main

import "fmt"

type bank_account struct {
	acc_num int
	name    string
	balance float64
	deposit float64
}

var b1, b2 bank_account

func update_balance() {
	b1.balance += b1.deposit
	b2.balance += b2.deposit
}

func main() {

	fmt.Println("\nENTER THE DETAILS OF CUSTOMER 1 :")
	fmt.Print("ACCOUNT NUMBER : ")
	fmt.Scanln(&b1.acc_num)
	fmt.Print("ACCOUNT HOLDER NAME : ")
	fmt.Scanln(&b1.name)
	fmt.Print("ACCOUNT BALANCE : ")
	fmt.Scanln(&b1.balance)
	fmt.Print("DEPOSIT AMOUNT : ")
	fmt.Scanln(&b1.deposit)

	fmt.Println("\nENTER THE DETAILS OF CUSTOMER 2 :")
	fmt.Print("ACCOUNT NUMBER : ")
	fmt.Scanln(&b2.acc_num)
	fmt.Print("ACCOUNT HOLDER NAME : ")
	fmt.Scanln(&b2.name)
	fmt.Print("ACCOUNT BALANCE : ")
	fmt.Scanln(&b2.balance)
	fmt.Print("DEPOSIT AMOUNT : ")
	fmt.Scanln(&b2.deposit)

	update_balance()
	fmt.Printf("\n\nACCOUNT BALANCE OF CUSTOMER 1 : RS.%.2f\n", b1.balance)
	fmt.Printf("\nACCOUNT BALANCE OF CUSTOMER 2 : RS.%.2f\n", b2.balance)
}
