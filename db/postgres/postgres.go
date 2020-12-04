package postgres

import (
	"log"
	"movie/db"
	"movie/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id varchar(36) PRIMARY KEY,
	first_name varchar(100),
	last_name varchar(100),
	email varchar(100),
	password varchar(100)
);
`
var _ db.DB = &SqlxDB{}

type SqlxDB struct {
	conn *sqlx.DB
}

func New(dns string) db.DB {
	conn, err := sqlx.Connect("postgres", dns)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("postgres conn", conn.Ping())
	log.Println(conn.MustExec(schema).RowsAffected())

	return &SqlxDB{conn}
}

// Users
func (db *SqlxDB) AddUser(u *model.User) (*model.User, error) {
	query := `INSERT INTO users (id, first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5);`
	db.conn.MustExec(query, u.ID, u.FirstName, u.LastName, u.Email, u.Password)
	return u, nil
}
func (db *SqlxDB) DeleteUser(uuid string) error {
	db.conn.Query("DELETE FROM users WHERE id = $1", uuid)
	return nil
}
func (db *SqlxDB) GetUsers() (map[string]*model.User, error) {
	var users []model.User
	db.conn.Select(&users, "SELECT * FROM users")
	finalres := make(map[string]*model.User)
	for _, u := range users {
		finalres[u.ID] = &u
	}
	return finalres, nil
}
func (db *SqlxDB) GetUserByUUID(uuid string) (*model.User, error) {
	var u model.User
	db.conn.Get(&u, "SELECT * FROM users WHERE id = $1", uuid)
	return &u, nil
}

func (db *SqlxDB) UpdateUser(uuid string, data map[string]interface{}) (*model.User, error) {
	// var user model.User
	return nil, nil
}

func (db *SqlxDB) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	db.conn.Get(&u, "SELECT * FROM users WHERE email = $1", email)
	return &u, nil
}

func (SqlxDB) AddMedia(u *model.Media) (*model.Media, error) {
	return nil, nil
}
func (SqlxDB) DeleteMedia(uuid string) error {
	return nil
}
func (SqlxDB) GetMedias() (map[string]*model.Media, error) {
	return nil, nil
}
func (SqlxDB) GetMediaByUUID(uuid string) (*model.Media, error) {
	return nil, nil
}
func (SqlxDB) UpdateMedia(uuid string, data map[string]interface{}) (*model.Media, error) {
	return nil, nil
}

func (SqlxDB) AddActor(u *model.Actor) (*model.Actor, error) {
	return nil, nil
}
func (SqlxDB) DeleteActor(uuid string) error {
	return nil
}
func (SqlxDB) GetActors() (map[string]*model.Actor, error) {
	return nil, nil
}
func (SqlxDB) GetActorByUUID(uuid string) (*model.Actor, error) {
	return nil, nil
}
func (SqlxDB) UpdateActor(uuid string, data map[string]interface{}) (*model.Actor, error) {
	return nil, nil
}

func (SqlxDB) AddMovie(u *model.Movie) (*model.Movie, error) {
	return nil, nil
}
func (SqlxDB) GetMovies() (map[string]*model.Movie, error) {
	return nil, nil
}
func (SqlxDB) DeleteMovie(uuid string) error {
	return nil
}

func (SqlxDB) GetMovieByUUID(uuid string) (*model.Movie, error) {
	return nil, nil
}

func (SqlxDB) UpdateMovie(uuid string, data map[string]interface{}) (*model.Movie, error) {
	return nil, nil
}
