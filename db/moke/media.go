package moke

import "movie/model"

func (db *MokeDB) AddMedia(u *model.Media) (*model.Media, error) {
	db.Medias[u.ID] = u
	return u, nil
}

func (db *MokeDB) DeleteMedia(uuid string) error {
	delete(db.Medias, uuid)
	return nil
}

func (db *MokeDB) GetMedias() (map[string]*model.Media, error) {
	return db.Medias, nil
}

func (db *MokeDB) GetMediaByUUID(uuid string) (*model.Media, error) {
	return db.Medias[uuid], nil
}

func (db *MokeDB) UpdateMedia(uuid string, data map[string]interface{}) (*model.Media, error) {
	u, err := db.GetMediaByUUID(uuid)
	if err != nil {
		return nil, err
	}

	if v, ok := data["title"]; ok {
		u.Title = v.(string)
	}
	if v, ok := data["uri"]; ok {
		u.URI = v.(string)
	}
	if v, ok := data["kind"]; ok {
		u.Kind = model.MediaKindType(v.(string))
	}

	return u, nil
}
