package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var key *rsa.PrivateKey
var tpl *template.Template

func init() {
	var err error
	key, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/public", getPublicKey)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tpl.Execute(w, nil)
	case "POST":
		enc := r.FormValue("msg")
		dec, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, []byte(enc), []byte{})
		if err != nil {
			fmt.Println("enc", enc)
		} else {
			fmt.Println(string(dec))
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getPublicKey(w http.ResponseWriter, r *http.Request) {

	/*{   //this is an example jwk key, other key types are Uint8Array objects
	        kty: "RSA",
	        e: ,
	        n: ,
	        alg: "RSA-OAEP-256",
	        ext: true,
		}*/
	jason := map[string]interface{}{}
	jason["kty"] = "RSA"
	jason["e"] = base64.RawURLEncoding.EncodeToString([]byte(strconv.Itoa(key.PublicKey.E)))
	jason["n"] = base64.RawURLEncoding.EncodeToString(key.PublicKey.N.Bytes())
	jason["alg"] = "RSA-OAEP-256"
	jason["ext"] = true
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(jason)
}
