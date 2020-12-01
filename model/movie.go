package model

import (
	"time"

	"github.com/google/uuid"
)

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
