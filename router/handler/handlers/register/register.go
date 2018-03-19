// Package register
// 15 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package register

import (
	"net/http"

	"github.com/KaiserGald/gmscreen/router/handler/handle"
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
	}
	if r.Method == "POST" {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		if err := r.ParseForm(); err != nil {
			route.Log().Debug.Log("ParseForm() err: %v", err)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		route.Log().Debug.Log("Username: %v\nPassword: %v\n", username, password)
	}
}

// Route returns a pointer to the route
func Route() *handle.Route {
	return &route
}
