package main

import (
	"bufio"
	"fmt"
	"os"
)

type book struct {
	author    string
	title     string
	publisher string
	price     float64
	year      int
	edition   string
}

func update(b *book, price float64, year int, edition string) {
	b.price = price
	b.year = year
	b.edition = edition
}

func display(b *book) {
	fmt.Printf("%s, \"%s\", %s, %s, %d\n", b.author, b.title, b.publisher, b.edition, b.year)
}

func total_price(b *book) float64 {
	var copies float64
	fmt.Print("\nENTER THE NUMBER OF COPIES : ")
	fmt.Scanln(&copies)
	total := copies * b.price
	if b.year < 2020 {
		total -= (total * 0.2)
	}
	return total
}

func read(b *book) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nENTER AUTHOR'S NAME : ")
	author, _ := reader.ReadString('\n')
	b.author = author

	fmt.Print("\nENTER BOOK TITLE : ")
	title, _ := reader.ReadString('\n')
	b.title = title

	fmt.Print("\nENTER PUBLISHER'S NAME : ")
	publisher, _ := reader.ReadString('\n')
	b.publisher = publisher

	fmt.Print("\nENTER BOOK PRICE : ")
	fmt.Scanln(&b.price)

	fmt.Print("\nENTER YEAR OF PUBLICATION : ")
	fmt.Scanln(&b.year)

	fmt.Print("\nENTER EDITION : ")
	edition, _ := reader.ReadString('\n')
	b.edition = edition
}

func main() {
	var b1, b2 book
	var ch int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nENTER THE DETAILS OF BOOK 1 :")
	read(&b1)
	fmt.Print("\nDO YOU WISH TO UPDATE MAKE UPDATES (0 = NO, 1 = YES) : ")
	fmt.Scanln(&ch)
	if ch == 1 {
		var price float64
		var year int
		fmt.Print("\nENTER BOOK PRICE : ")
		fmt.Scanln(&price)
		fmt.Print("\nENTER YEAR OF PUBLICATION : ")
		fmt.Scanln(&year)
		fmt.Print("\nENTER EDITION : ")
		edition, _ := reader.ReadString('\n')
		update(&b1, price, year, edition)
	}

	total := total_price(&b1)
	display(&b1)
	fmt.Println("\nTOTAL PRICE : ", total)

	fmt.Println("\nENTER THE DETAILS OF BOOK 2 :")
	read(&b2)

	fmt.Print("\nDO YOU WISH TO UPDATE MAKE UPDATES : (0 = NO, 1 = YES)")
	fmt.Scanln(&ch)
	if ch == 1 {
		var price float64
		var year int
		var edition string
		fmt.Print("\nENTER BOOK PRICE : ")
		fmt.Scanln(price)
		fmt.Print("\nENTER YEAR OF PUBLICATION : ")
		fmt.Scanln(year)
		fmt.Print("\nENTER EDITION : ")
		fmt.Scanln(edition)
		update(&b2, price, year, edition)
	}

	total = total_price(&b2)
	display(&b2)
	fmt.Println("\nTOTAL PRICE : ", total)
}
