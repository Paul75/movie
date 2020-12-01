package main

import (
	"time"

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

// Movie represent a movie
type Movie struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"desc"`
	Actors       []Actor   `json:"actors,omitempty"`
	Medias       []Media   `json:"medias,omitempty"`
	CreationDate time.Time `json:"creation_date"`
}

func NewMovie(title, desc string, create time.Time) *Movie {
	return &Movie{
		ID:           uuid.New().String(),
		Title:        title,
		Description:  desc,
		CreationDate: create,
	}
}

type Actor struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}

type MediaKind int

const (
	MediaPic MediaKind = iota + 1
	MediaMovie
)

type Media struct {
	ID    string
	Title string
	URI   string
	Kind  MediaKind
}
