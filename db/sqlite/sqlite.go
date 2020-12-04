package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"movie/db"
	"movie/model"
)

var _ db.DB = &DBSQLite{}

type DBSQLite struct {
	db *gorm.DB
}

func New() db.DB {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})

	return &DBSQLite{db}
}

// Users
func (sqldb *DBSQLite) AddUser(u *model.User) (*model.User, error) {
	return u, sqldb.db.Create(u).Error
}
func (sqldb *DBSQLite) DeleteUser(uuid string) error {
	return sqldb.db.Where("id = ?", uuid).Delete(&model.User{}).Error
}
func (sqldb *DBSQLite) GetUsers() (map[string]*model.User, error) {
	var users []model.User
	sqldb.db.Find(&users)
	finalres := make(map[string]*model.User)
	for _, u := range users {
		finalres[u.ID] = &u
	}
	return finalres, nil
}
func (sqldb *DBSQLite) GetUserByUUID(uuid string) (*model.User, error) {
	var u model.User
	return &u, sqldb.db.Model(&u).Where("id = ?", uuid).First(&u).Error
}

func (sqldb *DBSQLite) UpdateUser(uuid string, data map[string]interface{}) (*model.User, error) {
	var user model.User

	return &user, sqldb.db.Model(&user).Where("id = ?", uuid).Updates(data).Error
}

func (sqldb *DBSQLite) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	return &user, sqldb.db.Where("email = ?", email).First(&user).Error
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
