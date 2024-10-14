package users

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	_ "embed"
)

// Obviously not threadsafe

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users = []User{
	{ID: 1, Name: "Alice", Age: 18},
	{ID: 2, Name: "Bob", Age: 65},
}

//go:embed user.tpl
var UserTempl string

//go:embed users.tpl
var UsersTempl string

func AddHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", usersAll)
	mux.HandleFunc("POST /users", usersAdd)
	mux.HandleFunc("GET /users/{id}", usersOne)
}
func usersAll(w http.ResponseWriter, r *http.Request) {
	renderUsers(w, Users)
}

func usersOne(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	idx := id - 1
	if idx < 0 || idx >= len(Users) {
		http.NotFound(w, r)
		return
	}

	renderUsers(w, []User{Users[idx]})
}

func renderUsers(w http.ResponseWriter, u []User) {
	tpl := template.Must(template.New("users").Parse(UsersTempl))
	tpl = template.Must(tpl.Parse(UserTempl))

	if err := tpl.ExecuteTemplate(w, "users", u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func usersAdd(w http.ResponseWriter, r *http.Request) {

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.ID = len(Users) + 1
	Users = append(Users, u)

	w.WriteHeader(http.StatusCreated)
	renderUsers(w, []User{u})
}
