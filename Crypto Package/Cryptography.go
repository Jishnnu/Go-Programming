package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

func main() {
	key := "Thisissecret123usedtoencrypttext" // len should be 32
	plaintext := "This is a secret"
	ciphertext := EncryptAES([]byte(key), plaintext)
	fmt.Println("PLAINTEXT : ", plaintext)
	fmt.Println("ENCRYPTED TEXT : ", ciphertext)
	DecryptAES([]byte(key), ciphertext)
}

func EncryptAES(key []byte, plaintext string) string {
	c, _ := aes.NewCipher(key)
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))
	return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ciphertext string) {
	cipher, _ := hex.DecodeString(ciphertext)
	c, _ := aes.NewCipher(key)
	plaintext := make([]byte, len(ciphertext))
	c.Decrypt(plaintext, cipher)
	s := string(plaintext[:])
	fmt.Println("DECRYPTED TEXT : ", s)
}
