package entities

type (
	Shot struct {
		Order    int      `json:"order,omitempty"`
		Played   bool     `json:"played,omitempty"`
		ShotData ShotData `json:"shotData"`
		ShotId   string   `json:"shotId,omitempty"`
		Status   string   `json:"status,omitempty"`
	}

	ShotData struct {
		CoverUri string   `json:"coverUri,omitempty"`
		MdsUrl   string   `json:"mdsUrl,omitempty"`
		ShotText string   `json:"shotText,omitempty"`
		ShotType ShotType `json:"shotType"`
	}

	ShotType struct {
		Id    string `json:"id,omitempty"`
		Title string `json:"title,omitempty"`
	}

	ShotEvent struct {
		EventId string `json:"eventId,omitempty"`
		Shots   []Shot `json:"shots,omitempty"`
	}
)
