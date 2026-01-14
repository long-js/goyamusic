package entities

type (
	Album struct {
		OriginalReleaseYear      interface{}   `json:"originalRelease_Year,omitempty"`
		Error                    string        `json:"error,omitempty"`
		Title                    string        `json:"title,omitempty"`
		Version                  string        `json:"version,omitempty"`
		CoverUri                 string        `json:"coverUri,omitempty"`
		ContentWarning           string        `json:"contentWarning,omitempty"`
		Genre                    string        `json:"genre,omitempty"`
		TextColor                string        `json:"textColor,omitempty"`
		ShortDescription         string        `json:"shortDescription,omitempty"`
		Description              string        `json:"description,omitempty"`
		MetaType                 string        `json:"metaType,omitempty"`
		StorageDir               string        `json:"storageDir,omitempty"`
		OgImage                  string        `json:"ogImage,omitempty"`
		ReleaseDate              string        `json:"releaseDate,omitempty"`
		Type                     string        `json:"type,omitempty"`
		StartDate                string        `json:"startDate,omitempty"`
		Buy                      []interface{} `json:"buy,omitempty"`
		Prerolls                 []interface{} `json:"prerolls,omitempty"`
		Bests                    []int         `json:"bests,omitempty"`
		Duplicates               []Album       `json:"duplicates,omitempty"`
		Albums                   []Album       `json:"albums,omitempty"`
		Volumes                  [][]Track     `json:"volumes,omitempty"`
		Artists                  []Artist      `json:"artists,omitempty"`
		Labels                   []Label       `json:"labels,omitempty"`
		Regions                  []string      `json:"regions,omitempty"`
		AvailableRegions         []string      `json:"availableRegions,omitempty"`
		AvailableForOptions      []string      `json:"availableForOptions,omitempty"`
		Year                     int           `json:"year,omitempty"`
		Id                       int           `json:"id,omitempty"`
		TrackCount               int           `json:"trackCount,omitempty"`
		DurationMs               int           `json:"durationMs,omitempty"`
		LikesCount               int           `json:"likesCount,omitempty"`
		Available                bool          `json:"available,omitempty"`
		AvailableForPremiumUsers bool          `json:"availableForPremiumUsers,omitempty"`
		IsPremiere               bool          `json:"isPremiere,omitempty"`
		IsBanner                 bool          `json:"isBanner,omitempty"`
		Recent                   bool          `json:"recent,omitempty"`
		VeryImportant            bool          `json:"veryImportant,omitempty"`
		AvailableForMobile       bool          `json:"availableForMobile,omitempty"`
		AvailablePartially       bool          `json:"availablePartially,omitempty"`
		AvailableAsRbt           bool          `json:"availableAsRbt,omitempty"`
		LyricsAvailable          bool          `json:"lyricsAvailable,omitempty"`
		RememberPosition         bool          `json:"rememberPosition,omitempty"`
		Explicit                 bool          `json:"explicit,omitempty"`
		TrackPosition            TrackPosition `json:"trackPosition"`
		Deprecation              Deprecation   `json:"deprecation"`
	}

	AlbumEvent struct {
		Album  Album `json:"album"`
		Tracks Track `json:"tracks"`
	}

	AlbumsResp struct {
		ResponseBase
		Result []Album `json:"result"`
	}
)
