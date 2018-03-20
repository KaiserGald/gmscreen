// Package models
// 19 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package models

import "github.com/KaiserGald/logger"

var log *logger.Logger

// Init initializes the models package
func Init(lg *logger.Logger) {
	log = lg
}
