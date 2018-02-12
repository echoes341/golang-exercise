package main

import (
	"net/http"

	"github.com/echoes341/golang-exercise/go-web-dev/042/tests/06/starting-code/controllers"
	"github.com/echoes341/golang-exercise/go-web-dev/042/tests/06/starting-code/models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(map[string]*models.User{})
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
