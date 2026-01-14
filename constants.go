package goyamusic

import (
	"time"

	"github.com/long-js/goyamusic/internal"
)

const (
	defaultTimeout = 5 * time.Second
)

type ObjType string

var (
	languages = internal.Set[string]{
		"ru": struct{}{},
		"en": struct{}{},
		"uz": struct{}{},
	}
	validObjectTypes = internal.Set[ObjType]{
		"artist":   struct{}{},
		"album":    struct{}{},
		"playlist": struct{}{},
		"track":    struct{}{},
	}
	validAcoountFields = internal.Set[string]{
		"userMusicVisibility":       struct{}{},
		"userSocialVisibility":      struct{}{},
		"theme":                     struct{}{},
		"volumePercents":            struct{}{},
		"lastFmScrobblingEnabled":   struct{}{},
		"shuffleEnabled":            struct{}{},
		"facebookScrobblingEnabled": struct{}{},
		"addNewTrackOnPlaylistTop":  struct{}{},
		"rbtDisabled":               struct{}{},
		"promosDisabled":            struct{}{},
		"autoPlayRadio":             struct{}{},
		"syncQueueEnabled":          struct{}{},
		"adsDisabled":               struct{}{},
		"diskEnabled":               struct{}{},
		"showDiskTracksInLibrary":   struct{}{},
	}
)

var AndroidKey = []byte("p93jhgh689SBReK6ghtw62")
