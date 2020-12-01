package model

import "github.com/google/uuid"

// User represent the customer.
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func NewUser(fn, ln, email, pwd string) *User {
	return &User{
		ID:        uuid.New().String(),
		FirstName: fn,
		LastName:  ln,
		Email:     email,
		Password:  pwd,
	}
}
