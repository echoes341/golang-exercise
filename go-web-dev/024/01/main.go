package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpg", dogImg)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "foo ran\n")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("dog.gohtml"))
	tpl.Execute(w, `<h1>This is from</h1>`)
}

func dogImg(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpg")
}
