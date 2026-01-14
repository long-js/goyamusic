package entities

type Feed struct {
	Today              string              `json:"today,omitempty"`
	NextRevision       string              `json:"nextRevision,omitempty"`
	CanGetMoreEvents   bool                `json:"canGetMoreEvents,omitempty"`
	Pumpkin            bool                `json:"pumpkin,omitempty"`
	IsWizardPassed     bool                `json:"isWizardPassed,omitempty"`
	GeneratedPlaylists []GeneratedPlaylist `json:"generatedPlaylists,omitempty"`
	Headlines          []interface{}       `json:"headlines,omitempty"`
	Days               []Day               `json:"days,omitempty"`
}
