package moke

import "movie/model"

func (db *MokeDB) AddUser(u *model.User) (*model.User, error) {
	db.Users[u.ID] = u
	return u, nil
}

func (db *MokeDB) DeleteUser(uuid string) error {
	delete(db.Users, uuid)
	return nil
}

func (db *MokeDB) GetUsers() (map[string]*model.User, error) {
	return db.Users, nil
}

func (db *MokeDB) GetUserByUUID(uuid string) (*model.User, error) {
	return db.Users[uuid], nil
}

func (db *MokeDB) UpdateUser(uuid string, data map[string]interface{}) (*model.User, error) {
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

	return u, nil
}
