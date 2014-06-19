package model

// import (
// 	"errors"
// 	"fmt"
// 	"time"

// 	"code.google.com/p/go.crypto/bcrypt"
// )

// var (
// 	ErrInvalidUserName = errors.New("Invalid User Name")
// 	ErrInvalidPassword = errors.New("Invalid Password")
// 	ErrInvalidEmail    = errors.New("Invalid Email Address")
// )

// // Gravatar URL pattern
// var GravatarPattern = "https://gravatar.com/avatar/%s?s=%v&d=identicon"

// // Simple regular expression used to verify that an email
// // address matches the expected standard format.
// var RegexpEmail = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

// type User struct {
// 	ID         int64      `meddler:"id,pk"             json:"id"`
// 	Email      string     `meddler:"email"             json:"email"`
// 	Password   string     `meddler:"password"          json:"-"`
// 	Token      string     `meddler:"token"             json:"-"`
// 	Name       string     `meddler:"name"              json:"name"`
// 	Gravatar   string     `meddler:"gravatar"          json:"gravatar"`
// 	Created    time.Time  `meddler:"created,utctime"   json:"created"`
// 	Updated    time.Time  `meddler:"updated,utctime"   json:"updated"`
// 	Admin      bool       `medder:"admin"              json:"-"`
// }

// // Creates a new User from the given Name and Email.
// func NewUser(name, email string) *User {
// 	user := User{}
// 	user.Name = name
// 	user.Token = createToken()
// 	user.SetEmail(email)
// 	return &user
// }