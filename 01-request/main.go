package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type responder int

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func (r responder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.html", req.Form)
}

func main() {
	var hdn responder
	err := http.ListenAndServe(":50174", hdn)
	if err != nil {
		log.Println(err)
	}
}
