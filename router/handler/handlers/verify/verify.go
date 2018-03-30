// Package verify
// 28 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package verify

import (
	"encoding/json"
	"net/http"

	"github.com/KaiserGald/gmscreen/db"
	"github.com/KaiserGald/gmscreen/models"
	"github.com/KaiserGald/gmscreen/router/handler/handle"
)

// Route is the route that will be used
var route handle.Route

func init() {
	route = handle.Route{}
	route.SetName("/verify")
	route.SetHandler(handleFunc)
}

// handleFunc is the actual handler for the function
func handleFunc(w http.ResponseWriter, r *http.Request) {
	route.Log().Debug.Log("Handling Route '/verify' Method '%v'.", r.Method)

	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			route.Log().Debug.Log("ParseForm() err: %v", err)
		}
		email := r.FormValue("email")
		token := r.FormValue("token")
		route.Log().Debug.Log(email)
		s, err := db.Connect()
		if err != nil {
			route.Log().Error.Log("Error connecting to datatase.")
		} else {
			route.Log().Info.Log("Connected to database.")
		}

		msg := handle.EmailVerificationMessage{
			TokenValid:   true,
			TokenExpired: false,
		}
		u, err := models.GetUserByEmail(email, s)
		if err != nil {
			route.Log().Error.Log("Error getting user: %v", err)
		}

		if u.EmailVerified {
			w.WriteHeader(http.StatusBadRequest)
			msg.EmailVerified = true
		} else {
			msg.EmailVerified = false
			if err = u.VerifyUserEmail(token, s); err != nil {
				route.Log().Error.Log("Email verification request invalid: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				if err.Error() == "authentication tokens do not match" {
					msg.TokenValid = false
				}
				if err.Error() == "token has already expired" {
					msg.TokenExpired = true
				}
			}
		}
		w.Header().Add("Access-Control-Allow-Origin", "*")
		jmsg, err := json.Marshal(msg)
		if err != nil {
			route.Log().Error.Log("Error marshalling json response: %v", err)
		}
		w.Write(jmsg)
	}

}

// Route returns a pointer to the route
func Route() *handle.Route {
	return &route
}
