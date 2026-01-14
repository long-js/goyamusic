package entities

type (
	Track struct {
		Title                          string             `json:"title,omitempty"`
		OgImage                        string             `json:"ogImage,omitempty"`
		Type                           string             `json:"type,omitempty"`
		CoverUri                       string             `json:"coverUri,omitempty"`
		StorageDir                     string             `json:"storageDir,omitempty"`
		Error                          string             `json:"error,omitempty"`
		State                          string             `json:"state,omitempty"`
		DesiredVisibility              string             `json:"desiredVisibility,omitempty"`
		Filename                       string             `json:"filename,omitempty"`
		ContentWarning                 string             `json:"contentWarning,omitempty"`
		Version                        string             `json:"version,omitempty"`
		BackgroundVideoUri             string             `json:"backgroundVideoUri,omitempty"`
		ShortDescription               string             `json:"shortDescription,omitempty"`
		TrackSharingFlag               string             `json:"trackSharingFlag,omitempty"`
		TrackSource                    string             `json:"trackSource,omitempty"`
		Id                             string             `json:"id,omitempty"`
		RealId                         string             `json:"realId,omitempty"`
		Substituted                    *Track             `json:"substituted,omitempty"`
		MatchedTrack                   *Track             `json:"matchedTrack,omitempty"`
		Artists                        []Artist           `json:"artists,omitempty"`
		Albums                         []Album            `json:"albums,omitempty"`
		PoetryLoverMatches             []PoetryLoverMatch `json:"poetryLoverMatches,omitempty"`
		Regions                        []string           `json:"regions,omitempty"`
		AvailableForOptions            []string           `json:"availableForOptions,omitempty"`
		DurationMs                     int                `json:"durationMs,omitempty"`
		FileSize                       int                `json:"fileSize,omitempty"`
		PreviewDurationMs              int                `json:"previewDurationMs,omitempty"`
		Available                      bool               `json:"available,omitempty"`
		AvailableForPremiumUsers       bool               `json:"availableForPremiumUsers,omitempty"`
		LyricsAvailable                bool               `json:"lyricsAvailable,omitempty"`
		Best                           bool               `json:"best,omitempty"`
		CanPublish                     bool               `json:"canPublish,omitempty"`
		AvailableAsRbt                 bool               `json:"availableAsRbt,omitempty"`
		Explicit                       bool               `json:"explicit,omitempty"`
		AvailableFullWithoutPermission bool               `json:"availableFullWithoutPermission,omitempty"`
		RememberPosition               bool               `json:"rememberPosition,omitempty"`
		IsSuitableForChildren          bool               `json:"isSuitableForChildren,omitempty"`
		Major                          Major              `json:"major"`
		Normalization                  Normalization      `json:"normalization"`
		UserInfo                       User               `json:"userInfo"`
		MetaData                       MetaData           `json:"metaData"`
		R128                           R128               `json:"r128"`
		LyricsInfo                     LyricsInfo         `json:"lyricsInfo"`
	}

	TrackShort struct {
		Timestamp     string `json:"timestamp,omitempty"`
		AlbumId       string `json:"albumId,omitempty"`
		Id            int    `json:"id,omitempty"`
		PlayCount     int    `json:"playCount,omitempty"`
		OriginalIndex int    `json:"originalIndex,omitempty"`
		Recent        bool   `json:"recent,omitempty"`
		TrackFull     *Track `json:"track,omitempty"`
	}

	TrackId struct {
		From_   string `json:"from_,omitempty"`
		Id      int    `json:"id,omitempty"`
		TrackId int    `json:"trackId,omitempty"`
		AlbumId int    `json:"albumId,omitempty"`
	}

	TrackIdTs struct {
		Id        string `json:"id,omitempty"`
		AlbumId   string `json:"albumId,omitempty"`
		Timestamp string `json:"timestamp,omitempty"`
	}

	TrackPosition struct {
		Volume int `json:"volume,omitempty"`
		Index  int `json:"index,omitempty"`
	}

	TracksSimilar struct {
		SimilarTracks []Track `json:"similarTracks,omitempty"`
		Track         *Track  `json:"track"`
	}

	TrackLyrics struct {
		DownloadUrl     string      `json:"downloadUrl,omitempty"`
		ExternalLyricId string      `json:"externalLyricId,omitempty"`
		Writers         []string    `json:"writers,omitempty"`
		LyricId         int         `json:"lyricId,omitempty"`
		Major           LyricsMajor `json:"major"`
	}

	TrackLikesResp struct {
		ResponseBase
		Result struct {
			Library Library `json:"library,omitempty"`
		} `json:"result"`
	}
)

type (
	TrackLosslessInfo struct {
		Quality string `json:"quality"`
		Codec   string `json:"codec"`
		Url     string `json:"url"`
		Bitrate int    `json:"bitrate"`
	}

	TrackLosslessInfoResp struct {
		ResponseBase
		Result struct {
			Info TrackLosslessInfo `json:"downloadInfo"`
		} `json:"result"`
	}
)
