package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var correct, total uint

	f, err := os.Open("../problems.csv")
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(os.Stdin)

	cF := csv.NewReader(f)
	rc, err := cF.Read()
	for err != io.EOF {
		total++
		fmt.Printf("%s: ", rc[0])
		s.Scan()
		a := s.Text()

		if a == rc[1] {
			correct++
		}
		rc, err = cF.Read()
	}

	fmt.Println()
	fmt.Println("Correct, Total")
	fmt.Println(correct, total)
}
