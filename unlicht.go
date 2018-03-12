// Package main
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/KaiserGald/gmscreen/daemon"
	"github.com/KaiserGald/logger"
)

var port int
var verbose bool
var quiet bool
var color bool
var log *logger.Logger
var logLevel int

func processCLI() *daemon.Config {
	log = logger.New()
	cfg := &daemon.Config{}

	processFlags(cfg)

	configureDaemon(cfg)

	return cfg
}

func processFlags(cfg *daemon.Config) {

	flag.BoolVar(&color, "color", false, "Colors the server output.")
	flag.BoolVar(&color, "c", false, "Colors the server output.")

	flag.IntVar(&port, "port", 8080, "sets listen port for server")
	flag.IntVar(&port, "p", 8080, "sets listen port for server")

	flag.BoolVar(&verbose, "verbose", false, "sets server log to Verbose mode")
	flag.BoolVar(&verbose, "v", false, "sets server log to Verbose mode")

	flag.BoolVar(&quiet, "quiet", false, "sets server log to Verbose mode")
	flag.BoolVar(&quiet, "q", false, "sets server log to Verbose mode")

	flag.Parse()

}

func configureDaemon(cfg *daemon.Config) {

	if verbose && quiet {
		log.Notice.Log("Can't start server in both verbose and quiet mode. Only use one. Defaulting to normal output mode...")
	}

	env := os.Getenv("RUN_ENV")
	log.SetLogLevel(logger.All)
	log.Debug.Log(env)
	log.Info.Log("Configuring server daemon...")
	if env == "DEV" {
		cfg.DevMode = true
		cfg.ListenSpec = ":" + strconv.Itoa(port)
		log.Debug.Log("Started in dev mode.")
	} else if env == "PROD" {
		cfg.DevMode = false
		cfg.ListenSpec = ":" + os.Getenv("PORT")
		log.SetLogLevel(logger.Normal)
		log.Debug.Log("Started in production mode.")
	}

	if verbose && !quiet {
		log.SetLogLevel(logger.Verbose)
		log.Debug.Log("Started in verbose mode.")
	}

	if quiet && !verbose {
		log.SetLogLevel(logger.ErrorsOnly)
		log.Debug.Log("Started in quiet mode.")
	}
	if color {
		log.Debug.Log("Started in colored output mode.")
	}
	log.ShowColor(color)
}

func main() {
	cfg := processCLI()
	log.Notice.Log("Starting app '%v'!", os.Getenv("BINARY_NAME"))

	log.Info.Log("Starting server daemon...")
	if err := daemon.Run(cfg, log); err != nil {
		log.Error.Log("Error in main(): %v", err)
	}
}
