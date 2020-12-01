package moke

import "movie/model"

func (db *MokeDB) AddActor(u *model.Actor) (*model.Actor, error) {
	db.Actors[u.ID] = u
	return u, nil
}

func (db *MokeDB) DeleteActor(uuid string) error {
	delete(db.Actors, uuid)
	return nil
}

func (db *MokeDB) GetActors() (map[string]*model.Actor, error) {
	return db.Actors, nil
}

func (db *MokeDB) GetActorByUUID(uuid string) (*model.Actor, error) {
	return db.Actors[uuid], nil
}

func (db *MokeDB) UpdateActor(uuid string, data map[string]interface{}) (*model.Actor, error) {
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
