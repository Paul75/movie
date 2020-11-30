package main

import "time"

type MokeDB struct {
	Users  map[string]*User
	Movies map[string]*Movie
	Medias map[string]*Media
	Actors map[string]*Actor
}

func NewMokeDB() *MokeDB {
	var db MokeDB
	db.Users = make(map[string]*User)
	db.Movies = make(map[string]*Movie)
	db.Medias = make(map[string]*Media)
	db.Actors = make(map[string]*Actor)

	m := NewMovie("Star Wars", "lasjdl;ksfj", time.Now())
	db.Movies[m.ID] = m

	return &db
}
