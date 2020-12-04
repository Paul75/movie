package moke

import (
	"time"

	"movie/db"
	"movie/model"
)

var _ db.DB = &MokeDB{}

type MokeDB struct {
	Users  map[string]*model.User
	Movies map[string]*model.Movie
	Medias map[string]*model.Media
	Actors map[string]*model.Actor
}

func NewMokeDB() db.DB {
	var db MokeDB
	db.Users = make(map[string]*model.User)
	db.Movies = make(map[string]*model.Movie)
	db.Medias = make(map[string]*model.Media)
	db.Actors = make(map[string]*model.Actor)

	m := model.NewMovie("Star Wars", "lasjdl;ksfj", time.Now())
	db.Movies[m.ID] = m

	return &db
}
