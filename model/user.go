package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

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

func (u *User) UnmarshalJSON(b []byte) error {
	aux := struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}{}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	u.ID = aux.ID
	u.FirstName = aux.FirstName
	u.LastName = aux.LastName
	u.Email = aux.Email
	u.Password = HashPass(aux.Password)

	return nil
}

func HashPass(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (u User) MarshalJSON() ([]byte, error) {
	aux := struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}

	return json.Marshal(aux)
}
