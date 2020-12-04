package sqlite

import (
	"movie/db"
	"movie/model"
)

var _ db.DB = &DBSQLite{}

type DBSQLite struct {
}

func New() db.DB {
	return &DBSQLite{}
}

// Users
func (DBSQLite) AddUser(u *model.User) (*model.User, error) {
	return nil, nil
}
func (DBSQLite) DeleteUser(uuid string) error {
	return nil
}
func (DBSQLite) GetUsers() (map[string]*model.User, error) {
	return nil, nil
}
func (DBSQLite) GetUserByUUID(uuid string) (*model.User, error) {
	return nil, nil
}

func (DBSQLite) UpdateUser(uuid string, data map[string]interface{}) (*model.User, error) {
	return nil, nil
}

func (DBSQLite) GetUserByEmail(email string) (*model.User, error) {
	return nil, nil
}

func (DBSQLite) AddMedia(u *model.Media) (*model.Media, error) {
	return nil, nil
}
func (DBSQLite) DeleteMedia(uuid string) error {
	return nil
}
func (DBSQLite) GetMedias() (map[string]*model.Media, error) {
	return nil, nil
}
func (DBSQLite) GetMediaByUUID(uuid string) (*model.Media, error) {
	return nil, nil
}
func (DBSQLite) UpdateMedia(uuid string, data map[string]interface{}) (*model.Media, error) {
	return nil, nil
}

func (DBSQLite) AddActor(u *model.Actor) (*model.Actor, error) {
	return nil, nil
}
func (DBSQLite) DeleteActor(uuid string) error {
	return nil
}
func (DBSQLite) GetActors() (map[string]*model.Actor, error) {
	return nil, nil
}
func (DBSQLite) GetActorByUUID(uuid string) (*model.Actor, error) {
	return nil, nil
}
func (DBSQLite) UpdateActor(uuid string, data map[string]interface{}) (*model.Actor, error) {
	return nil, nil
}

func (DBSQLite) AddMovie(u *model.Movie) (*model.Movie, error) {
	return nil, nil
}
func (DBSQLite) GetMovies() (map[string]*model.Movie, error) {
	return nil, nil
}
func (DBSQLite) DeleteMovie(uuid string) error {
	return nil
}

func (DBSQLite) GetMovieByUUID(uuid string) (*model.Movie, error) {
	return nil, nil
}

func (DBSQLite) UpdateMovie(uuid string, data map[string]interface{}) (*model.Movie, error) {
	return nil, nil
}
