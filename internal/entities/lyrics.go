package entities

type (
	Lyrics struct {
		Lyrics          string `json:"lyrics,omitempty"`
		FullLyrics      string `json:"fullLyrics,omitempty"`
		TextLanguage    string `json:"textLanguage,omitempty"`
		Url             string `json:"url,omitempty"`
		Id              int    `json:"id,omitempty"`
		HasRights       bool   `json:"hasRights,omitempty"`
		ShowTranslation bool   `json:"showTranslation,omitempty"`
	}

	LyricsMajor struct {
		Name       string `json:"name,omitempty"`
		PrettyName string `json:"prettyName,omitempty"`
		Id         int    `json:"id,omitempty"`
	}

	LyricsInfo struct {
		HasAvailableSyncLyrics bool `json:"hasAvailableSyncLyrics,omitempty"`
		HasAvailableTextLyrics bool `json:"hasAvailableTextLyrics,omitempty"`
	}
)
