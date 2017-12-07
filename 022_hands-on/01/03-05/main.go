package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func check(e error) {
	if e != nil {
		log.Fatalln("Something went wrong!")
	}
}

func main() {
	http.HandleFunc("/", defaultH)
	http.HandleFunc("/dog/", dogH)
	http.HandleFunc("/me/", meH)

	http.ListenAndServe(":8080", nil) //nil to use the default ServeMux
}

func defaultH(w http.ResponseWriter, req *http.Request) {
	check(tpl.ExecuteTemplate(w, "index.gohtml", nil))
}

func dogH(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	dogName := req.Form.Get("dog")
	check(tpl.ExecuteTemplate(w, "dog.gohtml", dogName))
}
func meH(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	check(tpl.ExecuteTemplate(w, "me.gohtml", req.Form["name"][0]))
}
