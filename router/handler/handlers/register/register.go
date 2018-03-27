// Package register
// 15 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package register

import (
	"encoding/json"
	"net/http"

	"github.com/KaiserGald/gmscreen/db"
	"github.com/KaiserGald/gmscreen/models"
	"github.com/KaiserGald/gmscreen/router/handler/handle"
	"golang.org/x/crypto/bcrypt"
)

// Route is the route that will be used
var route handle.Route

func init() {
	route = handle.Route{}
	route.SetName("/register")
	route.SetHandler(handleFunc)
}

// handleFunc is the actual handler for the function
func handleFunc(w http.ResponseWriter, r *http.Request) {
	route.Log().Debug.Log("Handling Route '/register'.")
	route.Log().Debug.Log(r.Method)
	switch r.Method {
	case "OPTIONS":
		route.PreflightHandler(w, r)

	case "GET":
		if err := r.ParseForm(); err != nil {
			route.Log().Debug.Log("ParseForm() err: %v", err)
		}
		username := r.FormValue("username")
		email := r.FormValue("email")
		s, err := db.Connect()
		if err != nil {
			route.Log().Error.Log("Error connecting to database.")
		} else {
			route.Log().Info.Log("Connected to database.")
		}
		msg := handle.ValidUserDataMessage{
			Username: false,
			Email:    false,
		}
		_, err = models.GetUserByName(username, s)
		if err == nil {
			msg.Username = true
		}
		_, err = models.GetUserByEmail(email, s)
		if err == nil {
			msg.Email = true
		}
		if msg.Username || msg.Email {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Add("Access-Control-Allow-Origin", "*")
		}
		if msg.Username {
			route.Log().Info.Log("Username already exists.")
		}
		if msg.Email {
			route.Log().Info.Log("Email already exists.")
		}
		jmsg, err := json.Marshal(msg)
		if err != nil {
			route.Log().Error.Log("Error marshalling json response: %v", err)
		}
		w.Write(jmsg)

	case "POST":
		w.Header().Add("Access-Control-Allow-Origin", "*")
		if err := r.ParseForm(); err != nil {
			route.Log().Debug.Log("ParseForm() err: %v", err)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		route.Log().Debug.Log("Username: %v\tPassword: %v\tEmail: %v", username, password, email)
		hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			route.Log().Debug.Log("Error hashing password: %v", err)
		}
		route.Log().Debug.Log("Hashed Password: %v", string(hp))
		s, err := db.Connect()
		if err != nil {
			route.Log().Error.Log("Error connecting to database.")
		} else {
			route.Log().Info.Log("Connected to database.")
		}
		if err = models.CreateUser(username, email, string(hp), s); err != nil {
			switch err.Error() {
			case "username already exists":
				route.Log().Info.Log("Username already taken.")
			case "email already exists":
				route.Log().Info.Log("Email already taken.")
			default:
				route.Log().Info.Log(err.Error())
			}
		}
		if err != nil {
			route.Log().Error.Log("Error connecting to database.")
		}
		u, err := models.GetUserByEmail(email, s)
		if err != nil {
			route.Log().Error.Log("Error getting user data: %v", err)
		}
		route.Log().Debug.Log("%v", u)
		err = route.Mailer.SendVerificationMail("Validate Your Email Address", "validate.html", []string{email}, u)
		if err != nil {
			route.Log().Error.Log("Error sending validation email: %v", err)
		}
	}
}

// Route returns a pointer to the route
func Route() *handle.Route {
	return &route
}
