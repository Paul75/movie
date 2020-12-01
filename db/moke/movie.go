package moke

import "movie/model"

func (db *MokeDB) AddMovie(u *model.Movie) (*model.Movie, error) {
	db.Movies[u.ID] = u
	return u, nil
}

func (db *MokeDB) GetMovies() (map[string]*model.Movie, error) {
	return db.Movies, nil
}
