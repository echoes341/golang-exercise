package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"os"
)

func main() {

	inFile, err := os.Open("encrypted.aes")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Insert a password")
		s.Scan()
		p := []byte(s.Text())
		block, err := aes.NewCipher(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Not valid password, try again")
			continue
		}
		var iv [aes.BlockSize]byte
		stream := cipher.NewOFB(block, iv[:])
		reader := &cipher.StreamReader{S: stream, R: inFile}
		fmt.Println("Trying to decrypt...")
		fmt.Println()
		fmt.Println("====================================")
		n, _ := io.Copy(os.Stdout, reader)
		if n == 0 {
			fmt.Println("ERROR: not valid password")
		}
		fmt.Println()
		fmt.Println("====================================")
	}
}
