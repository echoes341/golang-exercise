// hmac test
package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	// read from stdInput
	s := bufio.NewScanner(os.Stdin)
	var encrypted []byte
	for {
		fmt.Println("Insert a some text here")
		s.Scan()
		t := s.Text()

		mac := hmac.New(sha256.New, []byte("whatafeeling"))
		mac.Write([]byte(t))
		if len(encrypted) != 0 {
			current := mac.Sum(nil)
			if hmac.Equal(current, encrypted) {
				fmt.Println("Logged in ^^")
			}
		} else {
			encrypted = mac.Sum(nil)
			fmt.Println(encrypted)
		}
	}
}
