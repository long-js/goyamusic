package entities

type (
	Account struct {
		Now                  string          `json:"now,omitempty"`
		Login                string          `json:"login,omitempty"`
		FullName             string          `json:"fullName,omitempty"`
		SecondName           string          `json:"secondName,omitempty"`
		FirstName            string          `json:"firstName,omitempty"`
		DisplayName          string          `json:"displayName,omitempty"`
		Birthday             string          `json:"birthday,omitempty"`
		RegisteredAt         string          `json:"registeredAt,omitempty"`
		PassportPhones       []PassportPhone `json:"passportPhones,omitempty"`
		Region               int             `json:"region,omitempty"`
		Uid                  uint32          `json:"uid,omitempty"`
		HostedUser           bool            `json:"hostedUser,omitempty"`
		ServiceAvailable     bool            `json:"serviceAvailable,omitempty"`
		HasInfoForAppMetrica bool            `json:"hasInfoForAppMetrica,omitempty"`
		Child                bool            `json:"child,omitempty"`
	}

	AccountSettings struct {
		Modified                  string `json:"modified,omitempty"`
		UserMusicVisibility       string `json:"userMusicVisibility,omitempty"`
		UserSocialVisibility      string `json:"userSocialVisibility,omitempty"`
		Theme                     string `json:"theme,omitempty"`
		Uid                       uint32 `json:"uid,omitempty"`
		VolumePercents            int    `json:"volumePercents,omitempty"`
		LastFmScrobblingEnabled   bool   `json:"lastFmScrobblingEnabled,omitempty"`
		ShuffleEnabled            bool   `json:"shuffleEnabled,omitempty"`
		FacebookScrobblingEnabled bool   `json:"facebookScrobblingEnabled,omitempty"`
		AddNewTrackOnPlaylistTop  bool   `json:"addNewTrackOnPlaylistTop,omitempty"`
		RbtDisabled               bool   `json:"rbtDisabled,omitempty"`
		PromosDisabled            bool   `json:"promosDisabled,omitempty"`
		AutoPlayRadio             bool   `json:"autoPlayRadio,omitempty"`
		SyncQueueEnabled          bool   `json:"syncQueueEnabled,omitempty"`
		AdsDisabled               bool   `json:"adsDisabled,omitempty"`
		DiskEnabled               bool   `json:"diskEnabled,omitempty"`
		ShowDiskTracksInLibrary   bool   `json:"showDiskTracksInLibrary,omitempty"`
	}

	AccountSettingsResp struct {
		ResponseBase
		Result AccountSettings
	}

	AccountStatus struct {
		Advertisement  string       `json:"advertisement,omitempty"`
		DefaultEmail   string       `json:"defaultEmail,omitempty"`
		Userhash       string       `json:"userhash,omitempty"`
		SkipsPerHour   int          `json:"skipsPerHour,omitempty"`
		SubeditorLevel int          `json:"subeditorLevel,omitempty"`
		CacheLimit     int          `json:"cacheLimit,omitempty"`
		PremiumRegion  int          `json:"premiumRegion,omitempty"`
		Experiment     int          `json:"experiment,omitempty"`
		PretrialActive bool         `json:"pretrialActive,omitempty"`
		StationExists  bool         `json:"stationExists,omitempty"`
		Subeditor      bool         `json:"subeditor,omitempty"`
		Account        Account      `json:"account"`
		Permissions    Permissions  `json:"permissions"`
		Subscription   Subscription `json:"subscription"`
		Plus           Plus         `json:"plus"`
		StationData    StationData  `json:"stationData"`
		BarBelow       Alert        `json:"barBelow"`
	}

	AccountStatusResp struct {
		ResponseBase
		Result AccountStatus `json:"result"`
	}
)
