package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./templates/index.gohtml"))
}

func main() {
	pics := http.StripPrefix("/resources", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/", handle)
	http.Handle("/resources/", pics)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
