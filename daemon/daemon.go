// Package daemon
// 16 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package daemon

import (
	"crypto/tls"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KaiserGald/gmscreen/router"
	"github.com/KaiserGald/gmscreen/services/com/comhandler"
	"github.com/KaiserGald/gmscreen/services/com/comserver"
	"github.com/KaiserGald/logger"
)

var log *logger.Logger

// Config contains configuration information for the server
type Config struct {
	ListenSpec string
	DevMode    bool
}

// Run starts up the server daemon
func Run(cfg *Config, lg *logger.Logger) error {
	log = lg

	cer, err := tls.LoadX509KeyPair("/go/src/github.com/KaiserGald/gmscreen/data/certs/server.crt", "/go/src/github.com/KaiserGald/gmscreen/data/certs/server.key")
	if err != nil {
		log.Error.Log("Error loading certificates: %v", err)
		return err
	}

	config := &tls.Config{
		Certificates:       []tls.Certificate{cer},
		InsecureSkipVerify: true,
	}
	log.Notice.Log("Starting HTTP listener on %s", cfg.ListenSpec)
	l, err := net.Listen("tcp", cfg.ListenSpec)
	if err != nil {
		log.Error.Log("Error creating listener: %v", err)
		return err
	}

	err = router.Start(l, config, log)
	if err != nil {
		log.Error.Log("Error Starting Router.")
		return err
	}
	comserver.Start(log)
	comhandler.Start(log)

	log.Notice.Log("Server up and running.")

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch

	log.Debug.Log("Got signal: %v, exiting...", s)
	time.Sleep(2 * time.Second)
}
