package controllers

import (
	"errors"
	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
	"github.com/jeypc/go-auth/models"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

var UserModel = models.NewUserModel()

type UserInput struct {
	Username, Password string
}

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			temp, err := template.ParseFiles("auth/view/index.html")

			if err != nil {
				panic(err)
				temp.Execute(w, nil)
			}

		}
	}

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

		var message error
		if user.Username == "" {
			message = errors.New("Username atau Password Salah")
		} else {
			ErrPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			if ErrPassword != nil {
				message = errors.New("Password Salah")
			}
		}

		if message != nil {

			data := map[string]interface{}{
				"error": message,
			}
			temp, _ := template.ParseFiles("auth/view/login.html")
			temp.Execute(w, data)
		} else {
			session, _ := config.Store.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["username"] = user.Username
			session.Values["password"] = user.Password

			session.Save(r, w)

			http.Redirect(w, r, "/", http.StatusSeeOther)

		}

	}
}
