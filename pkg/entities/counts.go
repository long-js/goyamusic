package entities

type Counts struct {
	Tracks       int `json:"tracks,omitempty"`
	DirectAlbums int `json:"directAlbums,omitempty"`
	AlsoAlbums   int `json:"alsoAlbums,omitempty"`
	AlsoTracks   int `json:"alsoTracks,omitempty"`
}
