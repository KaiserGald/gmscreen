// Package models
// 19 March 2018
// Code is licensed under the MIT License
// © 2018 Scott Isenberg

package models

import (
	"errors"
	"time"

	"github.com/KaiserGald/gmscreen/utils/token"
	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// User models user account data.
type User struct {
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	CreatedOn string `json:"createdOn" bson:"createdOn"`
	UserEmail
}

// UserEmail models user email data
type UserEmail struct {
	Email         string `json:"email" bson:"email"`
	EmailVerified bool   `json:"emailVerified" bson:"emailVerified"`
	EmailToken    string `json:"emailToken" bson:"emailToken"`
	TokenExpires  string `json:"tokenExpires" bson:"tokenExpires"`
}

// CreateUser creates a new user and inserts it into the database.
func CreateUser(un, em, pw string, s *mgo.Session) error {
	t, err := token.GenerateToken(em)
	if err != nil {
		return err
	}
	u := User{
		Username:  un,
		Password:  pw,
		CreatedOn: time.Now().Format(time.RFC3339),
		UserEmail: UserEmail{
			Email:         em,
			EmailVerified: false,
			EmailToken:    t,
			TokenExpires:  time.Now().Add(time.Hour * 6).Format(time.RFC3339),
		},
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

// VerifyUserEmail sets the EmailVerified field to true
func (u *User) VerifyUserEmail(token string, s *mgo.Session) error {
	if token == u.EmailToken {
		u.EmailVerified = true
	} else {
		return errors.New("authentication tokens do not match")
	}

	exp, err := time.Parse(time.RFC3339, u.TokenExpires)
	if err != nil {
		return err
	}
	if time.Now().After(exp) {
		u.EmailVerified = false
		return errors.New("token has already expired")
	}

	c := s.DB("gmscreen").C("User")
	colQ := bson.M{"username": u.Username}
	change := bson.M{"$set": bson.M{"useremail.emailVerified": true}}
	err = c.Update(colQ, change)
	if err != nil {
		return err
	}

	return nil
}

// NewEmailVerifyToken creates a new email verification token, resets the expiration time and updates the database. It then returns the token.
func (u *User) NewEmailVerifyToken(s *mgo.Session) (string, error) {
	t, err := token.GenerateToken(u.Email)
	if err != nil {
		return "", err
	}

	exp := time.Now().Add(time.Hour * 6).Unix()
	c := s.DB("gmscreen").C("User")
	colQ := bson.M{"username": u.Username}
	change := bson.M{"$set": bson.M{"emailToken": t, "tokenExpires": exp}}
	err = c.Update(colQ, change)
	if err != nil {
		return "", err
	}
	return t, nil
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
	if err := c.Find(bson.M{"useremail.email": em}).One(&res); err != nil {
		return res, err
	}
	return res, nil
}

// DeleteUser removes a user from the database.
func DeleteUser() {

}
