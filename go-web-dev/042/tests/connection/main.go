package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	_, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection estabilished")
}
