package controllers

import (
	"github.com/jeypc/go-auth/entities"
	"github.com/jeypc/go-auth/models"
	"html/template"
	"net/http"
)

var UserModel = models.NewUserModel()

type UserInput struct {
	Username, Password string
}

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("auth/view/index.html")

	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, err := template.ParseFiles("auth/view/login.html")

		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		var user entities.User
		UserModel.Where(&user, "username", UserInput.Username)

		// if user.Username == ""
	}
}
