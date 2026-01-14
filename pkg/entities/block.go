package entities

type Block struct {
	Id          string                                     `json:"id,omitempty"`
	Type        string                                     `json:"type,omitempty"`
	TypeForFrom string                                     `json:"typeForFrom,omitempty"`
	Title       string                                     `json:"title,omitempty"`
	Description string                                     `json:"description,omitempty"`
	Entities    []BlockEntity                              `json:"entities,omitempty"`
	Data        map[PersonalPlaylistsData]PlayContextsData `json:"data,omitempty"`
}
