package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var correct, total uint
	var path string

	flag.StringVar(&path, "f", "../problems.csv", "CSV file containing problems")
	flag.Parse()

	f, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(os.Stdin)
	fmt.Println("Press enter to start the quiz")
	s.Scan()

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
	fmt.Println("Correct/Total")
	fmt.Printf("%d/%d\n", correct, total)
}
