package model

type MediaKind int

const (
	MediaPic MediaKind = iota + 1
	MediaMovie
)

type Media struct {
	ID    string
	Title string
	URI   string
	Kind  MediaKind
}
