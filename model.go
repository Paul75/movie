package main

import (
	"time"

	"github.com/google/uuid"
)

// User represent the customer.
type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// Movie represent a movie
type Movie struct {
	ID           string
	Title        string
	Description  string
	Actors       []Actor
	Medias       []Media
	CreationDate time.Time
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
	ID        string
	FirstName string
	LastName  string
	Biography string
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
