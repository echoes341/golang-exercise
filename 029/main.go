package main

import "html/template"
import "net/http"
import "log"
import "strconv"

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./templates/index.gohtml"))
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func main() {
	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}
func index(w http.ResponseWriter, req *http.Request) {
	var times int
	//Create cookie if !exists
	c, err := req.Cookie("times")
	if err == http.ErrNoCookie {
		times = 0
	} else {
		times, _ = strconv.Atoi(c.Value)
	}
	times++
	http.SetCookie(w, &http.Cookie{
		Name:  "times",
		Value: strconv.Itoa(times),
	})

	tpl.ExecuteTemplate(w, "index.gohtml", times)

}
