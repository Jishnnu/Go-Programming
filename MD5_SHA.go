package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	fmt.Println("--------------------- SMALL MESSAGE --------------------")
	message := []byte("Welcome to Go Programming. Today's lesson is Crypto Package and its methods.")
	fmt.Printf("MD5 : %x\n", md5.Sum(message))
	fmt.Printf("SHA-1 : %x\n", sha1.Sum(message))
	fmt.Printf("SHA-256 : %x\n", sha256.Sum256(message))
	fmt.Printf("SHA-512 : %x\n", sha512.Sum512(message))

	fmt.Println("\n--------------------- LARGE MESSAGE --------------------")
	message = []byte("As a tech-driven financial company, one of our biggest goals at Capital One is to use technology to easily, quickly, and directly engage with our customers. One important way a potential Capital One customer can engage with us is through our affiliate channel. Affiliate partners like Credit Sesame, CreditCards.com, and Bankrate have a special partnership with Capital One where they display available credit card options and guide potential customers to the credit cards that are right for them. For our affiliate partners, this is done through the Credit Offers API. We like to think that it helps people make smarter credit card choices. We also like to think it shows off our cutting edge use of technology.")
	fmt.Printf("MD5 : %x\n", md5.Sum(message))
	fmt.Printf("SHA-1 : %x\n", sha1.Sum(message))
	fmt.Printf("SHA-256 : %x\n", sha256.Sum256(message))
	fmt.Printf("SHA-512 : %x\n", sha512.Sum512(message))
}
