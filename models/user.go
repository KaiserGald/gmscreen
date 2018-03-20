// Package models
// 19 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package models

import (
	"errors"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// User models user account data.
type User struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// CreateUser creates a new user and inserts it into the database.
func CreateUser(un, em, pw string, s *mgo.Session) error {

	u := User{
		Username: un,
		Email:    em,
		Password: pw,
	}
	log.Debug.Log("%v", u)
	s.SetMode(mgo.Monotonic, true)

	c := s.DB("gmscreen").C("User")

	user, err := GetUserByName(u.Username, s)
	if user.Username != "" {
		return errors.New("username already exists")
	}
	if err != nil && err.Error() != "not found" {
		return err
	}

	user, err = GetUserByEmail(u.Email, s)
	if user.Email != "" {
		return errors.New("email already exists")
	}
	if err != nil && err.Error() != "not found" {
		return err
	}

	err = c.Insert(u)
	if err != nil {
		log.Error.Log("Error inserting user into database: %v", err)
		return err
	}
	return nil
}

// UpdateUserEmail updates the user's email with a new one.
func (u *User) UpdateUserEmail() {

}

// UpdateUserPassword updates the user's password with a new one.
func (u *User) UpdateUserPassword() {

}

// GetUserByName gets a user by name and returns a user object.
func GetUserByName(un string, s *mgo.Session) (User, error) {
	var res User
	c := s.DB("gmscreen").C("User")
	if err := c.Find(bson.M{"username": un}).One(&res); err != nil {
		return res, err
	}
	return res, nil
}

// GetUserByEmail gets a user by email and returns a user object.
func GetUserByEmail(em string, s *mgo.Session) (User, error) {
	var res User
	c := s.DB("gmscreen").C("User")
	if err := c.Find(bson.M{"email": em}).One(&res); err != nil {
		return res, err
	}
	return res, nil
}

// DeleteUser removes a user from the database.
func DeleteUser() {

}
