package moke

import (
	"movie/model"
	"time"
)

func (db *MokeDB) AddMovie(u *model.Movie) (*model.Movie, error) {
	db.Movies[u.ID] = u
	return u, nil
}

func (db *MokeDB) DeleteMovie(uuid string) error {
	delete(db.Movies, uuid)
	return nil
}
func (db *MokeDB) GetMovies() (map[string]*model.Movie, error) {
	return db.Movies, nil
}

func (db *MokeDB) GetMovieByUUID(uuid string) (*model.Movie, error) {
	return db.Movies[uuid], nil
}

func (db *MokeDB) UpdateMovie(uuid string, data map[string]interface{}) (*model.Movie, error) {
	m, err := db.GetMovieByUUID(uuid)
	if err != nil {
		return nil, err
	}

	if v, ok := data["title"]; ok {
		m.Title = v.(string)
	}
	if v, ok := data["desc"]; ok {
		m.Description = v.(string)
	}
	if v, ok := data["actors"]; ok {
		m.Actors = v.([]model.Actor)
	}
	if v, ok := data["medias"]; ok {
		m.Medias = v.([]model.Media)
	}
	if v, ok := data["creation_date"]; ok {
		t, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			return nil, err
		}
		m.CreationDate = t
	}

	return m, nil
}
