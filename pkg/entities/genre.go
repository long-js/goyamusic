package entities

type Genre struct {
	Id            string            `json:"id,omitempty"`
	Title         string            `json:"title,omitempty"`
	FullTitle     string            `json:"fullTitle,omitempty"`
	UrlPart       string            `json:"urlPart,omitempty"`
	Color         string            `json:"color,omitempty"`
	Weight        int               `json:"weight,omitempty"`
	ComposerTop   bool              `json:"composerTop,omitempty"`
	ShowInMenu    bool              `json:"showInMenu,omitempty"`
	HideInRegions interface{}       `json:"hideInRegions,omitempty"`
	ShowInRegions []interface{}     `json:"showInRegions,omitempty"`
	Titles        map[string]string `json:"titles,omitempty"`
	Images        Images            `json:"images"`
	RadioIcon     Icon              `json:"radioIcon"`
	SubGenres     []Genre           `json:"subGenres,omitempty"`
}
