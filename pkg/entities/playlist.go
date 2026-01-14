package entities

type (
	Playlist struct {
		Title                 string          `json:"title,omitempty"`
		Visibility            string          `json:"visibility,omitempty"`
		UrlPart               string          `json:"urlPart,omitempty"`
		Created               string          `json:"created,omitempty"`
		Modified              string          `json:"modified,omitempty"`
		OgImage               string          `json:"ogImage,omitempty"`
		OgTitle               string          `json:"ogTitle,omitempty"`
		OgDescription         string          `json:"ogDescription,omitempty"`
		Image                 string          `json:"image,omitempty"`
		BackgroundColor       string          `json:"backgroundColor,omitempty"`
		TextColor             string          `json:"textColor,omitempty"`
		IdForFrom             string          `json:"idForFrom,omitempty"`
		DummyDescription      string          `json:"dummyDescription,omitempty"`
		DummyPageDescription  string          `json:"dummyPageDescription,omitempty"`
		GeneratedPlaylistType string          `json:"generatedPlaylistType,omitempty"`
		AnimatedCoverUri      string          `json:"animatedCoverUri,omitempty"`
		Description           string          `json:"description,omitempty"`
		DescriptionFormatted  string          `json:"descriptionFormatted,omitempty"`
		PlaylistUuid          string          `json:"playlistUuid,omitempty"`
		Type                  string          `json:"type,omitempty"`
		Tags                  []Tag           `json:"tags,omitempty"`
		Coauthors             []int           `json:"coauthors,omitempty"`
		TopArtist             []Artist        `json:"topArtist,omitempty"`
		RecentTracks          []TrackId       `json:"recentTracks,omitempty"`
		Tracks                []TrackShort    `json:"tracks,omitempty"`
		SimilarPlaylists      []Playlist      `json:"similarPlaylists,omitempty"`
		LastOwnerPlaylists    []Playlist      `json:"lastOwnerPlaylists,omitempty"`
		Prerolls              []interface{}   `json:"prerolls,omitempty"`
		Uid                   uint32          `json:"uid,omitempty"`
		Kind                  int             `json:"kind,omitempty"`
		TrackCount            int             `json:"trackCount,omitempty"`
		Revision              int             `json:"revision,omitempty"`
		Snapshot              int             `json:"snapshot,omitempty"`
		DurationMs            int             `json:"durationMs,omitempty"`
		MetrikaId             int             `json:"metrikaId,omitempty"`
		LikesCount            int             `json:"likesCount,omitempty"`
		Collective            bool            `json:"collective,omitempty"`
		Available             bool            `json:"available,omitempty"`
		IsBanner              bool            `json:"isBanner,omitempty"`
		IsPremiere            bool            `json:"isPremiere,omitempty"`
		EverPlayed            bool            `json:"everPlayed,omitempty"`
		Ready                 bool            `json:"ready,omitempty"`
		Owner                 User            `json:"owner"`
		Cover                 Cover           `json:"cover"`
		MadeFor               MadeFor         `json:"madeFor"`
		PlayCounter           PlayCounter     `json:"playCounter"`
		PlaylistAbsence       PlaylistAbsence `json:"playlistAbsence"`
		CoverWithoutText      Cover           `json:"coverWithoutText"`
		Contest               Contest         `json:"contest"`
		DummyCover            Cover           `json:"dummyCover"`
		DummyRolloverCover    Cover           `json:"dummyRolloverCover"`
		OgData                OpenGraphData   `json:"ogData"`
		Branding              Brand           `json:"branding"`
		CustomWave            CustomWave      `json:"customWave"`
		Pager                 Pager           `json:"pager"`
	}

	PlaylistId struct {
		Uid  uint32 `json:"uid,omitempty"`
		Kind int    `json:"kind,omitempty"`
	}

	PlaylistAbsence struct {
		Reason string `json:"reason,omitempty"`
		Kind   int    `json:"kind,omitempty"`
	}

	PlaylistRecommendation struct {
		BatchId string  `json:"batchId,omitempty"`
		Tracks  []Track `json:"tracks,omitempty"`
	}

	PlaylistLikesResp struct {
		ResponseBase
		Result []Playlist `json:"result"`
	}

	PlaylistResp struct {
		ResponseBase
		Result *Playlist `json:"result"`
	}
)
