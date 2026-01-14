package entities

type BlockEntity struct {
	Data interface{} `json:"data,omitempty"` // TODO: пока не понятно что нужно воткнуть
	Id   string      `json:"id,omitempty"`
	Type string      `json:"type,omitempty"`
}
