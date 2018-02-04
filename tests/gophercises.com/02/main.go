package main

import (
	"log"
	"net/http"
)

type yamlPaths struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func main() {
	/*paths := map[string]string{
		"/gg": "https://www.google.com",
		"/fb": "https://www.facebook.com",
	}
	http.Handle("/", MapHandler(paths, http.NotFoundHandler()))*/
	yaml := `
- path: /gg
  url: https://www.google.com
- path: /fb
  url: https://www.facebook.com
`
	hdlr, err := YAMLHandler([]byte(yaml), http.NotFoundHandler())
	if err != nil {
		log.Fatalln("Not a valid yaml.", err)
	}
	http.HandleFunc("/", hdlr)
	http.ListenAndServe(":8080", nil)
}
