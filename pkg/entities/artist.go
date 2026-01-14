package entities

import "strconv"

type (
	Artist struct {
		Name                 string
		OgImage              string
		OpImage              string
		HandMadeDescription  string
		EnWikipediaLink      string
		InitDate             string
		EndDate              string
		YaMoneyId            string
		Id                   ArtistId
		NoPicturesFromSearch interface{}
		FullNames            interface{}
		Aliases              interface{}
		Genres               []string
		Regions              []string
		Decomposed           []string
		Countries            []string
		DbAliases            []string
		Links                []Link
		PopularTracks        []Track
		LikesCount           int
		Various              bool
		Composer             bool
		Available            bool
		TicketsAvailable     bool
		Cover                Cover
		Counts               Counts
		Ratings              Raitings
		Description          Description
	}

	ArtistId struct {
		Value int
	}

	// Deprecated
	ArtistTracks struct {
		Tracks []Track `json:"tracks"`
		Pager  Pager   `json:"pager"`
	}

	// Deprecated
	ArtistAlbums struct {
		Albums []Album `json:"albums,omitempty"`
		Pager  Pager   `json:"pager,omitempty"`
	}

	ArtistEvent struct {
		Tracks                      []Track  `json:"tracks,omitempty"`
		SimilarToArtistsFromHistory []Artist `json:"similarToArtistsFromHistory,omitempty"`
		Subscribed                  bool     `json:"subscribed,omitempty"`
		Artist                      Artist   `json:"artist"`
	}

	ArtistsResp struct {
		ResponseBase
		Result []Artist `json:"result"`
	}
)

func (a *ArtistId) UnmarshalJSON(b []byte) error {
	if b[0] == 34 {
		b = b[1 : len(b)-1]
	}
	v, err := strconv.ParseInt(string(b), 10, 32)
	(*a).Value = int(v)
	return err
}
