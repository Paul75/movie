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

func (db *MokeDB) AddUser(u *User) (*User, error) {
	db.Users[u.ID] = u
	return u, nil
}

func (db *MokeDB) DeleteUser(uuid string) error {
	delete(db.Users, uuid)
	return nil
}

func (db *MokeDB) GetUsers() (map[string]*User, error) {
	return db.Users, nil
}

func (db *MokeDB) GetUserByUUID(uuid string) (*User, error) {
	return db.Users[uuid], nil
}

func (db *MokeDB) UpdateUser(uuid string, data map[string]interface{}) (*User, error) {
	u, err := db.GetUserByUUID(uuid)
	if err != nil {
		return nil, err
	}

	if v, ok := data["first_name"]; ok {
		u.FirstName = v.(string)
	}
	if v, ok := data["last_name"]; ok {
		u.LastName = v.(string)
	}
	if v, ok := data["email"]; ok {
		u.Email = v.(string)
	}
	if v, ok := data["password"]; ok {
		u.Password = v.(string)
	}
	return u, nil
}

func (db *MokeDB) AddActor(u *Actor) (*Actor, error) {
	db.Actors[u.ID] = u
	return u, nil
}

func (db *MokeDB) DeleteActor(uuid string) error {
	delete(db.Actors, uuid)
	return nil
}

func (db *MokeDB) GetActors() (map[string]*Actor, error) {
	return db.Actors, nil
}

func (db *MokeDB) GetActorByUUID(uuid string) (*Actor, error) {
	return db.Actors[uuid], nil
}

func (db *MokeDB) UpdateActor(uuid string, data map[string]interface{}) (*Actor, error) {
	a, err := db.GetActorByUUID(uuid)
	if err != nil {
		return nil, err
	}

	if v, ok := data["first_name"]; ok {
		a.FirstName = v.(string)
	}
	if v, ok := data["last_name"]; ok {
		a.LastName = v.(string)
	}
	if v, ok := data["biography"]; ok {
		a.Biography = v.(string)
	}

	return a, nil
}
