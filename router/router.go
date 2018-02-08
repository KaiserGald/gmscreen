// Package ui
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package router

import (
	"net"
	"net/http"
	"time"

	"github.com/KaiserGald/logger"
	"github.com/KaiserGald/unlichtServer/router/handler"
)

// Start initializes and starts the ui router
func Start(listener net.Listener, log *logger.Logger) error {

	cacheResource := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			e := "\"" + r.URL.Path + "\""
			w.Header().Add("Etag", e)
			w.Header().Add("Cache-Control", "max-age=691200")
			h.ServeHTTP(w, r)
		}
	}

	log.Info.Log("Starting front-end.\n")

	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16,
	}
	err := handler.Start(log)
	if err != nil {
		log.Error.Log("Error starting route handler.")
		return err
	}
	handler.Handle()
	http.Handle("/img/", http.StripPrefix("/img", cacheResource(http.FileServer(http.Dir("/srv/unlichtServer/app/assets/img")))))
	http.Handle("/css/", http.StripPrefix("/css", cacheResource(http.FileServer(http.Dir("/srv/unlichtServer/app/assets/css")))))
	http.Handle("/js/", http.StripPrefix("/js", cacheResource(http.FileServer(http.Dir("/srv/unlichtServer/app/assets/js")))))
	go server.Serve(listener)

	log.Info.Log("Front-end up and running.")
	return nil
}
