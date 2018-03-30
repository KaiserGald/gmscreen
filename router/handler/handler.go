// Package handler
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package handler

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/KaiserGald/gmscreen/router/handler/handle"
	"github.com/KaiserGald/gmscreen/router/handler/handlers/index"
	"github.com/KaiserGald/gmscreen/router/handler/handlers/login"
	"github.com/KaiserGald/gmscreen/router/handler/handlers/register"
	"github.com/KaiserGald/gmscreen/router/handler/handlers/verify"
	"github.com/KaiserGald/logger"
)

var (
	routes []handle.Route
	log    *logger.Logger
)

// Start starts the handler
func Start(lg *logger.Logger) error {
	log = lg
	log.Debug.Log("Starting route handler.")
	index.Route().Init(log)
	err := Add(index.Route())
	if err != nil {
		return err
	}

	login.Route().Init(log)
	err = Add(login.Route())
	if err != nil {
		return err
	}

	register.Route().Init(log)
	err = Add(register.Route())
	if err != nil {
		return err
	}

	verify.Route().Init(log)
	err = Add(verify.Route())
	if err != nil {
		return err
	}

	return nil
}

// Add adds a new route to the handler
func Add(r *handle.Route) error {
	log.Debug.Log("Adding '%v' to handler.", r.Name())
	if compareRoute(r) {
		return errors.New("Route already exists")
	}
	routes = append(routes, *r)
	return nil
}

// Handle handles all the registered routes
func Handle() {
	log.Debug.Log("Route Handler Started.")
	for _, route := range routes {
		http.Handle(route.Name(), route.Handler())
	}
}

// compareRoute checks to see if a given route already exists
func compareRoute(r *handle.Route) bool {
	for i := range routes {
		if routes[i].Name() == r.Name() {
			return true
		}
	}
	return false
}

func validRouteName(s string) (bool, error) {
	r, err := regexp.MatchString("^/[a-zA-Z]+\\w+$", s)
	if err != nil {
		return r, err
	}
	return r, nil
}
