package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/echoes341/golang-exercise/go-web-dev/042/tests/08/controllers"
	"github.com/echoes341/golang-exercise/go-web-dev/042/tests/08/models"
	"github.com/julienschmidt/httprouter"
)

const fileName = "db.json"

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession(fileName), fileName)
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession(fileName string) map[string]models.User {
	db := make(map[string]models.User)
	f, err := os.Open(fileName)
	if err != nil {
		return db
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&db)
	return db
}
