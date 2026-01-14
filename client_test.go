package goyamusic

import (
	"context"
	"os"
	"testing"

	"github.com/long-js/goyamusic/pkg/entities"
)

var (
	client *Client
	acc    *entities.Account
)

func TestMain(m *testing.M) {
	var err error
	client, err = NewClient(os.Getenv("YAMUSIC_TOKEN"), "", "ru", nil)
	if err != nil {
		_, _ = os.Stderr.Write([]byte("token is absent:" + err.Error()))
		os.Exit(1)
	}

	ctx, cancelF := context.WithTimeout(context.Background(), defaultTimeout)
	if status, err := (*client).AccountStatus(ctx); err != nil {
		_, _ = os.Stderr.Write([]byte("account status fault:" + err.Error()))
		cancelF()
		os.Exit(2)
	} else {
		acc = &(*status).Account
	}
	cancelF()

	resCode := (*m).Run()
	os.Exit(resCode)
}

func TestAccountSettings(t *testing.T) {
	(*t).Skip()
	ctx, cancelF := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelF()

	if settings, err := (*client).AccountSettings(ctx); err != nil {
		(*t).Error("account settings fault:", err)
	} else {
		(*t).Logf("Account settings: %+v\n", settings)
	}
}

func TestAccountSettingsSet(t *testing.T) {
	(*t).Skip()
	ctx, cancelF := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelF()

	if settings, err := (*client).PutAccountSettings(ctx, "volumePercents", 40); err != nil {
		(*t).Error("account settings fault:", err)
	} else {
		(*t).Logf("Account settings: %+v\n", settings)
	}
}

func TestAccountMultiSettings(t *testing.T) {
	(*t).Skip()
	ctx, cancelF := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelF()

	data := map[string]interface{}{"foo": 40.1}
	if _, err := (*client).PutAccountSettingsMulti(ctx, data); err == nil {
		(*t).Error("account settings fault, need error")
	}

	data = map[string]interface{}{"volumePercents": 40}
	if settings, err := (*client).PutAccountSettingsMulti(ctx, data); err != nil {
		(*t).Error("account settings fault:", err)
	} else {
		(*t).Logf("Account settings: %+v\n", settings)
	}
}

func TestSettings(t *testing.T) {
	(*t).Skip()
	ctx, cancelF := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelF()

	if settings, err := (*client).Settings(ctx); err != nil {
		(*t).Error("settings fault:", err)
	} else {
		(*t).Logf("Settings: %+v\n", settings)
	}
}

func TestGetLikedObjects(t *testing.T) {
	(*t).Skip()
	ctx, cancelF := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelF()

	if res, err := (*client).Liked(ctx, (*acc).Uid, "artist"); err != nil {
		(*t).Error(err)
	} else {
		(*t).Logf("ARTISTS: %+v", res)
	}
	if res, err := (*client).Liked(ctx, (*acc).Uid, "album"); err != nil {
		(*t).Error(err)
	} else {
		(*t).Logf("ALBUMS: %+v", res)
	}
	if res, err := (*client).Liked(ctx, (*acc).Uid, "playlist"); err != nil {
		(*t).Error(err)
	} else {
		(*t).Logf("PLAYLISTS: %+v", res)
	}
	if res, err := (*client).Liked(ctx, (*acc).Uid, "track"); err != nil {
		(*t).Error(err)
	} else {
		(*t).Logf("TRACKS: %+v", res)
	}
}

func TestPlaylists(t *testing.T) {
	(*t).Skip()
	ctx, cancelF := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelF()

	if settings, err := (*client).Playlists(ctx, (*acc).Uid); err != nil {
		(*t).Error("playlists fault:", err)
	} else {
		(*t).Logf("playlists: %+v\n", settings)
	}
}

func TestPlaylist(t *testing.T) {
	(*t).Skip()
	ctx, cancelF := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelF()

	if tracks, err := (*client).PlaylistTracks(ctx, (*acc).Uid, "1006"); err != nil {
		(*t).Error("playlist fault:", err)
	} else {
		(*t).Logf("playlist tracks: %+v\n", tracks)
	}
}

func TestDownloadTrack(t *testing.T) {
	// (*t).Skip()
	ctxPl, cancelPl := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelPl()
	ctxDl, cancelDl := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancelDl()

	if tracks, err := (*client).PlaylistTracks(ctxPl, (*acc).Uid, "1006"); err != nil {
		(*t).Error("playlist fault:", err)
	} else if err = (*client).DownloadTrack(ctxDl, tracks[0].TrackFull); err != nil {
		(*t).Error("playlist fault:", err)
	} else {
		(*t).Log("track downloaded")
	}
}

// func TestPermissionAlerts(t *testing.T) {
// 	permissionAlerts, err := c.PermissionAlerts()
// 	if err != nil {
// 		log.Error("Unable to get permissionAlerts ", err)
// 	}
// 	log.Info("permissionAlerts: ", permissionAlerts)
//
// 	accountExperiments, err := c.AccountExperiments()
// 	if err != nil {
// 		log.Error("Unable to get accountExperiments ", err)
// 	}
// 	log.Info("accountExperiments: ", accountExperiments)
// }
//
// func TestPromo(t *testing.T) {
// 	promocode, err := c.ConsumePromocode("K6TS5BREDG", "en")
// 	if err != nil {
// 		log.Error("Unable to get promocode ", err)
// 	}
// 	log.Info("promocode: ", promocode)
// }
//
// func TestPermissionAlerts(t *testing.T) {
// 	likes, err := c.GetLikesTracks(strconv.Itoa(status.Account.Uid))
// 	if err != nil {
// 		log.Error("Unable to get likes ", err)
// 	}
// }
