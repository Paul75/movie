package db

import "movie/model"

type DB interface {
	DBUsers
	DBMedias
	DBActors
	DBMovies
}

type DBUsers interface {
	// Users
	AddUser(u *model.User) (*model.User, error)
	DeleteUser(uuid string) error
	GetUsers() (map[string]*model.User, error)
	GetUserByUUID(uuid string) (*model.User, error)
	UpdateUser(uuid string, data map[string]interface{}) (*model.User, error)
}

type DBMedias interface {
	// Medias
	AddMedia(u *model.Media) (*model.Media, error)
	DeleteMedia(uuid string) error
	GetMedias() (map[string]*model.Media, error)
	GetMediaByUUID(uuid string) (*model.Media, error)
	UpdateMedia(uuid string, data map[string]interface{}) (*model.Media, error)
}

type DBActors interface {
	// Actors
	AddActor(u *model.Actor) (*model.Actor, error)
	DeleteActor(uuid string) error
	GetActors() (map[string]*model.Actor, error)
	GetActorByUUID(uuid string) (*model.Actor, error)
	UpdateActor(uuid string, data map[string]interface{}) (*model.Actor, error)
}

type DBMovies interface {
	// Movies
	AddMovie(u *model.Movie) (*model.Movie, error)
	GetMovies() (map[string]*model.Movie, error)
}
