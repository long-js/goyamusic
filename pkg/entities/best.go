package entities

type Best struct {
	Result interface{} `json:"result,omitempty"` // TODO: [Union[Track, Artist, Album, Playlist, Video]]
	Type   string      `json:"type,omitempty"`
	Text   string      `json:"text,omitempty"`
}
