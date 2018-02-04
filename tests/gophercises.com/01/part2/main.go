package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var correct, total uint
	var path string

	flag.StringVar(&path, "f", "../problems.csv", "CSV file containing problems")
	duration := flag.Int("d", 30, "")
	flag.Parse()

	f, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(os.Stdin)
	fmt.Println("Press enter to start the quiz")
	s.Scan()

	timer := time.NewTimer(time.Second * time.Duration(*duration))
	cF := csv.NewReader(f)
	rc, err := cF.Read()
problemloop:
	for err != io.EOF {
		total++
		fmt.Printf("%s: ", rc[0])

		answer := make(chan string)

		go func() {
			s.Scan()
			answer <- s.Text()
		}()

		select {
		case <-timer.C:
			break problemloop
		case a := <-answer:
			if a == rc[1] {
				correct++
			}
			rc, err = cF.Read()
		}
	}

	fmt.Println()
	fmt.Printf("Correct/Total: %d/%d\n", correct, total)
}
