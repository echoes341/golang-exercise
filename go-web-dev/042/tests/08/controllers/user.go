package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/echoes341/golang-exercise/go-web-dev/042/tests/08/models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
)

type UserController struct {
	session  map[string]models.User
	nameFile string
}

func NewUserController(m map[string]models.User, fileName string) *UserController {
	return &UserController{m, fileName}
}

func (uc UserController) writeToFile() {
	f, err := os.OpenFile(uc.nameFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(f)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(uc.session)
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Retrieve user
	u := uc.session[id]

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	u.Id = uuid.NewV4().String()

	// store the user
	uc.session[u.Id] = u

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
	uc.writeToFile()
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	delete(uc.session, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
	uc.writeToFile()
}
