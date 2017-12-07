package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultH(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintln(w, "Welcome to the home page")
	if err != nil {
		log.Fatalln("Something went wrong")
	}
}
func dogH(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintln(w, "Elio is the name of my dog")
	if err != nil {
		log.Fatalln("Something went wrong")
	}
}
func meH(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintln(w, "Gianpaolo is my name")
	if err != nil {
		log.Fatalln("Something went wrong")
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(defaultH))
	http.Handle("/dog/", http.HandlerFunc(dogH))
	http.Handle("/me/", http.HandlerFunc(meH))

	http.ListenAndServe(":8080", nil) //nil to use the default ServeMux
}
