// Package login
// 15 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package login

import (
	"net/http"

	"github.com/KaiserGald/gmscreen/router/handler/handle"
	"github.com/satori/go.uuid"
)

// Route is the route that will be used
var route handle.Route

func init() {
	route = handle.Route{}
	route.SetName("/login")
	route.SetHandler(handleFunc)
}

// handleFunc is the actual handler for the function
func handleFunc(w http.ResponseWriter, r *http.Request) {
	route.Log().Debug.Log("Handling Route '/login' Method '%v'.", r.Method)
	route.Log().Debug.Log(r.Method)
	switch r.Method {
	case "OPTIONS":
		route.PreflightHandler(w, r)
	}
	if r.Method == "POST" {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		if err := r.ParseForm(); err != nil {
			route.Log().Debug.Log("ParseForm() err: %v", err)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		route.Log().Debug.Log("Username: %v\tPassword: %v", username, password)

		sID, err := uuid.NewV4()
		if err != nil {
			route.Log().Debug.Log("Error generating uuid: %v", err)
		}
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
}

// Route returns a pointer to the route
func Route() *handle.Route {
	return &route
}
