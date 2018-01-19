package main

import (
	"bufio"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	// generating random key
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	fmt.Println(key.PublicKey)
	s := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert a message here")
	s.Scan()
	t := []byte(s.Text())
	label := []byte("nothinginteresting")
	fmt.Println("Thank you. Encrypting message with random rsa keys...")
	enc, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &key.PublicKey, t, label)
	if err != nil {
		panic(err)
	}
	fmt.Println("This is your encoded message:")
	fmt.Println("======")
	fmt.Println(enc)
	fmt.Println("======")
	fmt.Println("Decrypting your message know")
	dec, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, enc, label)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dec))
}
