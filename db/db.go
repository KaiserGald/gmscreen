// Package models
// 19 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package db

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/KaiserGald/logger"
	mgo "github.com/globalsign/mgo"
)

var log *logger.Logger

// Config is a global config struct for the database
var Config config

// config contains the configuration for the database
type config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// Init initializes the database
func Init(lg *logger.Logger) error {
	log = lg
	Config = config{}
	path := os.Getenv("APPROOT")
	js, err := ioutil.ReadFile(path + "/data/conf/db.json")
	if err != nil {
		log.Error.Log("Error reading json file: %v", err)
		return err
	}

	json.Unmarshal(js, &Config)
	return nil
}

// URL returns the full URL of the database
func (c config) URL() string {
	return c.Host + ":" + c.Port
}

// Connect connects to the database and returns a database session
func Connect() (*mgo.Session, error) {
	s, err := mgo.Dial(Config.URL())
	if err != nil {
		log.Error.Log("Error connecting to database: %v", err)
		return &mgo.Session{}, err
	}
	return s, nil
}
