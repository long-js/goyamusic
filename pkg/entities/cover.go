package entities

type Cover struct {
	Type           string   `json:"type,omitempty"`
	Uri            string   `json:"uri,omitempty"`
	Dir            string   `json:"dir,omitempty"`
	Version        string   `json:"version,omitempty"`
	CopyrightName  string   `json:"copyrightName,omitempty"`
	CopyrightCline string   `json:"copyrightCline,omitempty"`
	Prefix         string   `json:"prefix,omitempty"`
	Error          string   `json:"error,omitempty"`
	ItemsUri       []string `json:"itemsUri,omitempty"`
	Custom         bool     `json:"custom,omitempty"`
	IsCustom       bool     `json:"isCustom,omitempty"`
}
