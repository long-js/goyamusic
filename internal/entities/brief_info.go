package entities

type BriefInfo struct {
	Concerts       interface{}  `json:"concerts,omitempty"`
	Albums         []Album      `json:"albums,omitempty"`
	Playlists      []Playlist   `json:"playlists,omitempty"`
	AlsoAlbums     []Album      `json:"alsoAlbums,omitempty"`
	LastReleaseIds []int        `json:"lastReleaseIds,omitempty"`
	LastReleases   []Album      `json:"lastReleases,omitempty"`
	PopularTracks  []Track      `json:"popularTracks,omitempty"`
	SimilarArtists []Artist     `json:"similarArtists,omitempty"`
	AllCovers      []Cover      `json:"allCovers,omitempty"`
	Videos         []Video      `json:"videos,omitempty"`
	Vinyls         []Vinyl      `json:"vinyls,omitempty"`
	PlaylistIds    []PlaylistId `json:"playlistIds,omitempty"`
	TracksInChart  []Chart      `json:"tracksInChart,omitempty"`
	HasPromotions  bool         `json:"hasPromotions,omitempty"`
	Artist         Artist       `json:"artist"`
}
