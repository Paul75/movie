package model

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

type MediaKind int

func (m *MediaKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*m = MediaKindType(s)
	return nil
}

func MediaKindType(str string) MediaKind {
	var m MediaKind
	switch strings.ToLower(str) {
	default:
		m = MediaDefault
	case "pic":
		m = MediaPic
	case "movie":
		m = MediaMovie
	}
	return m
}

func (m MediaKind) MarshalJSON() ([]byte, error) {
	var s string
	switch m {
	default:
		s = "unknown"
	case MediaPic:
		s = "pic"
	case MediaMovie:
		s = "movie"
	}
	return json.Marshal(s)
}

const (
	MediaPic MediaKind = iota + 1
	MediaMovie
	MediaDefault
)

type Media struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	URI   string    `json:"uri"`
	Kind  MediaKind `json:"kind"`
}

func NewMedia(title, uri string, kind MediaKind) *Media {
	return &Media{
		ID:    uuid.New().String(),
		Title: title,
		URI:   uri,
		Kind:  kind,
	}
}
