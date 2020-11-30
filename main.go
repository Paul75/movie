package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/movies", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

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
