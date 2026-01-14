package entities

type (
	Library struct {
		PlaylistUuid string      `json:"playlistUuid"`
		Uid          uint32      `json:"uid"`
		Revision     uint16      `json:"revision"`
		Artists      []Artist    `json:"artists,omitempty"`
		Albums       []Album     `json:"albums,omitempty"`
		Playlists    []Playlist  `json:"playlists,omitempty"`
		Tracks       []TrackIdTs `json:"tracks,omitempty"`
	}

	LibraryResp struct {
		ResponseBase
		Result struct {
			Library Library `json:"library"`
		} `json:"result"`
	}
)
