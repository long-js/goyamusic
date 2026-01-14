package entities

type AdParams struct {
	Id       string  `json:"id,omitempty"`
	Modified string  `json:"modified,omitempty"`
	Context  Context `json:"context"`
}
