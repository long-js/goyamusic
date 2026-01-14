package entities

type (
	Chart struct {
		Progress  string  `json:"progress,omitempty"`
		BgColor   string  `json:"bgColor,omitempty"`
		Position  int     `json:"position,omitempty"`
		Listeners int     `json:"listeners,omitempty"`
		Shift     int     `json:"shift,omitempty"`
		TrackId   TrackId `json:"trackId"`
	}

	ChartItem struct {
		Track Track `json:"track"`
		Chart Chart `json:"chart"`
	}

	ChartInfoMenuItem struct {
		Title    string `json:"title,omitempty"`
		Url      string `json:"url,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	}

	ChartInfoMenu struct {
		Items []ChartInfoMenuItem `json:"items"`
	}

	ChartInfo struct {
		Id               string        `json:"id,omitempty"`
		Type             string        `json:"type,omitempty"`
		TypeForFrom      string        `json:"typeForFrom,omitempty"`
		Title            string        `json:"title,omitempty"`
		ChartDescription string        `json:"chartDescription,omitempty"`
		Menu             ChartInfoMenu `json:"menu"`
		Chart            Playlist      `json:"chart"`
	}
)
