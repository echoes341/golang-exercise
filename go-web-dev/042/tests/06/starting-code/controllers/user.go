package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/echoes341/golang-exercise/go-web-dev/042/tests/06/starting-code/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
)

type UserController struct {
	db map[string]*models.User
}

func NewUserController(db map[string]*models.User) *UserController {
	return &UserController{db}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	if id != "" {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// composite literal
	var u *models.User
	var ok bool

	// Fetch user
	if u, ok = uc.db[id]; !ok {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u *models.User

	json.NewDecoder(r.Body).Decode(u)

	fmt.Printf("%v\n", u)

	// create bson ID
	id := uuid.NewV4().String()

	uc.db[id] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if id != "" {
		w.WriteHeader(404)
		return
	}

	// Delete user
	if _, ok := uc.db[id]; !ok {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
