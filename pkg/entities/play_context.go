package entities

type PlayContext struct {
	Context     string       `json:"context,omitempty"`
	ContextItem string       `json:"contextItem,omitempty"`
	Tracks      []TrackShort `json:"tracks,omitempty"`
}
