// Package handler
// 21 January 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package handle

// ValidUserDataMessage represents a json message about the validity of requested user data
type ValidUserDataMessage struct {
	Username bool
	Email    bool
}
