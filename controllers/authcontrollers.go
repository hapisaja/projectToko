package controllers

import (
	"fmt"
	"github.com/jeypc/go-auth/entities"
	"github.com/jeypc/go-auth/models"
	"html/template"
	"net/http"
)

type Input struct {
	username, password string
}

var UserModel = models.NewUserModel()

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		fmt.Println("Proses Login Mas")
		r.ParseForm()
		Input := &Input{
			username: r.Form.Get("username"),
			password: r.Form.Get("password"),
		}
	}
	var user entities.User
	UserModel.Where(&user, "username", Input.username)
}
