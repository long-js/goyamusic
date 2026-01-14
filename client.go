package goyamusic

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-json"
	ent "github.com/long-js/goyamusic/pkg/entities"
)

func NewClient(token, baseUrl, language string, httpClient *http.Client) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("unable to create client without token")
	}

	c := Client{baseUrl: baseUrl, language: language}

	if cwd, err := os.Getwd(); err != nil {
		return nil, fmt.Errorf("get working directory fault: %+v", err)
	} else {
		c.cacheDir = cwd + "/.cache"
	}

	headers := make(http.Header)
	headers["X-Yandex-Music-Client"] = []string{"YandexMusicAndroid/24023231"}
	headers["Authorization"] = []string{fmt.Sprintf("OAuth %s", token)}
	c.headers = headers

	if !languages.Has(language) {
		c.language = "en"
	} else {
		c.language = language
	}

	if baseUrl == "" {
		c.baseUrl = "https://api.music.yandex.net"
	} else {
		c.baseUrl = baseUrl
	}
	if httpClient == nil {
		c.httpClient = http.DefaultClient
	}
	c.device = "os=linux; os_version=; manufacturer=long-js; model=Yandex Music API; clid=; " +
		"device_id=random; uuid=random"

	return &c, nil
}

type Client struct {
	baseUrl    string
	language   string
	device     string
	cacheDir   string
	httpClient *http.Client
	headers    http.Header
}

func (c *Client) callAPI(ctx context.Context, method, path string, params url.Values, data, result interface{}) error {
	var (
		err    error
		bData  []byte
		req    *http.Request
		resp   *http.Response
		buffer *bytes.Buffer
		uri    string
	)
	if params != nil && len(params) > 0 {
		uri = fmt.Sprintf("%s/%s?%s", (*c).baseUrl, path, params.Encode())
	} else {
		uri = fmt.Sprintf("%s/%s", (*c).baseUrl, path)
	}
	// body
	if data != nil {
		if bData, err = json.Marshal(data); err != nil {
			return err
		}
		buffer = bytes.NewBuffer(bData)
	} else {
		buffer = &bytes.Buffer{}
	}

	if req, err = http.NewRequestWithContext(ctx, method, uri, buffer); err != nil {
		return fmt.Errorf("request creation fault: %+v", err)
	} else if len(bData) > 0 {
		(*c).headers["Content-Length"] = []string{strconv.FormatInt(int64(len(bData)), 10)}
	}
	req.Header = (*c).headers

	log.Printf("--> %+v", req.URL)
	resp, err = (*c).httpClient.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()

	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		var buf []byte

		if buf, err = io.ReadAll(resp.Body); err != nil {
			return err
		}
		log.Printf("<-- %s", buf)
		if err = json.Unmarshal(buf, result); err != nil {
			return err
		}
	case http.StatusUnauthorized:
		err = fmt.Errorf("req failed, check the token: %s", resp.Status)
	case http.StatusBadRequest:
		err = fmt.Errorf("req failed: %s", resp.Status)
	default:
		err = fmt.Errorf("req failed, unexpected responce code: %s", resp.Status)
	}
	return err
}

func (c *Client) download(uri, fpath string, out io.Writer) error {
	var (
		err  error
		file *os.File
		resp *http.Response
	)

	if fpath != "" {
		if file, err = os.Create(fpath); err != nil {
			return err
		}
		defer func() {
			_ = file.Close()
		}()
	}

	resp, err = (*c).httpClient.Get(uri)
	defer func() {
		if resp.Body != nil {
			_ = resp.Body.Close()
		}
	}()

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download fault: %s, %s", uri, resp.Status)
	}

	if fpath != "" {
		_, err = io.Copy(file, resp.Body)
	}
	if out != nil {
		_, err = io.Copy(out, resp.Body)
	}
	return err
}

func (c *Client) AccountStatus(ctx context.Context) (*ent.AccountStatus, error) {
	var result ent.AccountStatusResp

	if err := (*c).callAPI(ctx, http.MethodGet, "account/status", nil, nil, &result); err != nil {
		return nil, fmt.Errorf("account status fault: %+v", err)
	}
	if result.Error != "" {
		return nil, fmt.Errorf("account status fault: %s", result.Error)
	}
	return &result.Result, nil
}

func (c *Client) AccountSettings(ctx context.Context) (*ent.AccountSettings, error) {
	var result ent.AccountSettingsResp

	if err := (*c).callAPI(ctx, http.MethodGet, "account/settings", nil, nil, &result); err != nil {
		return nil, fmt.Errorf("account settings fault: %+v", err)
	}
	if result.Error != "" {
		return nil, fmt.Errorf("account settings fault: %s", result.Error)
	}
	return &result.Result, nil
}

func (c *Client) PutAccountSettings(ctx context.Context, key string, value interface{}) (*ent.AccountSettings, error) {
	var result ent.AccountSettingsResp

	if !validAcoountFields.Has(key) {
		return nil, fmt.Errorf("wrong setting data key")
	}

	err := (*c).callAPI(ctx, http.MethodPost, "account/settings", nil, map[string]interface{}{key: value}, &result)
	if err != nil {
		return nil, fmt.Errorf("account status fault: %+v", err)
	}
	if result.Error != "" {
		return nil, fmt.Errorf("account status fault: %s", result.Error)
	}

	return &result.Result, nil
}

func (c *Client) PutAccountSettingsMulti(ctx context.Context, data map[string]interface{}) (*ent.AccountSettings,
	error) {
	var result ent.AccountSettingsResp

	for k := range data {
		if !validAcoountFields.Has(k) {
			return nil, fmt.Errorf("wrong setting data key")
		}
	}

	if err := (*c).callAPI(ctx, http.MethodPost, "account/settings", nil, data, &result); err != nil {
		return nil, fmt.Errorf("account status fault: %+v", err)
	}
	if result.Error != "" {
		return nil, fmt.Errorf("account status fault: %s", result.Error)
	}

	return &result.Result, nil
}

/*
 * Likes
 */

func (c *Client) Settings(ctx context.Context) (*ent.Settings, error) {
	var result ent.SettingsResult

	if err := (*c).callAPI(ctx, http.MethodGet, "settings", nil, nil, &result); err != nil {
		return nil, fmt.Errorf("settings fault: %+v", err)
	}
	if result.Error != "" {
		return nil, fmt.Errorf("settings fault: %s", result.Error)
	}
	return &result.Result, nil
}

func (c *Client) Liked(ctx context.Context, userId uint32, objType ObjType) (interface{}, error) {
	if !validObjectTypes.Has(objType) {
		return nil, fmt.Errorf("wrong objType, supported: %v", validObjectTypes)
	}

	var (
		err  error
		data interface{}
		uri  = fmt.Sprintf("users/%d/likes/%ss?rich=true", userId, objType)
	)

	switch objType {
	case "artist":
		resp := ent.ArtistsResp{}
		err = (*c).callAPI(ctx, http.MethodGet, uri, nil, nil, &resp)
		if err == nil && resp.Error != "" {
			err = fmt.Errorf("get liked artists fault: %v", resp.Error)
		}
		data = resp.Result
	case "album":
		resp := ent.AlbumsResp{}
		err = (*c).callAPI(ctx, http.MethodGet, uri, nil, nil, &resp)
		if err == nil && resp.Error != "" {
			err = fmt.Errorf("get liked albums fault: %v", resp.Error)
		}
		data = resp.Result
	case "playlist":
		resp := ent.PlaylistLikesResp{}
		err = (*c).callAPI(ctx, http.MethodGet, uri, nil, nil, &resp)
		if err == nil && resp.Error != "" {
			err = fmt.Errorf("get liked playlists fault: %v", resp.Error)
		}
		data = resp.Result
	case "track":
		resp := ent.TrackLikesResp{}
		err = (*c).callAPI(ctx, http.MethodGet, uri, nil, nil, &resp)
		if err == nil && resp.Error != "" {
			err = fmt.Errorf("get liked tracks fault: %v", resp.Error)
		}
		data = resp.Result.Library.Tracks
	}

	if err != nil {
		return nil, fmt.Errorf("get liked %ss fault: %+v", objType, err)
	}
	return data, nil
}

func (c *Client) PutLikes(ctx context.Context, userId uint32, objType ObjType, ids []int) error {
	var result ent.LibraryResp

	if !validObjectTypes.Has(objType) {
		return fmt.Errorf("wrong objType, supported: %v", validObjectTypes)
	}

	uri := fmt.Sprintf("users/%d/likes/%ss/add-multiple", userId, objType)
	data := map[string][]int{fmt.Sprintf("%s-ids", objType): ids}

	if err := (*c).callAPI(ctx, http.MethodPost, uri, nil, data, &result); err != nil {
		return fmt.Errorf("get liked %s fault: %+v", objType, err)
	}
	if result.Error != "" {
		return fmt.Errorf("get liked %s fault: %s", objType, result.Error)
	}
	return nil
}

/*
 * Playlists
 */

func (c *Client) Playlists(ctx context.Context, userId uint32) ([]ent.Playlist, error) {
	var (
		result ent.PlaylistLikesResp
		path   = fmt.Sprintf("users/%d/playlists/list", userId)
	)
	if err := (*c).callAPI(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, fmt.Errorf("get playlists fault: %+v", err)
	}
	if result.Error != "" {
		return nil, fmt.Errorf("get playlists fault: %s", result.Error)
	}
	return result.Result, nil
}

func (c *Client) PlaylistTracks(ctx context.Context, userId uint32, playlistKind string) ([]ent.TrackShort, error) {
	var (
		result ent.PlaylistResp
		path   = fmt.Sprintf("users/%d/playlists/%s", userId, playlistKind)
	)
	if err := (*c).callAPI(ctx, http.MethodGet, path, nil, nil, &result); err != nil {
		return nil, fmt.Errorf("get playlists fault: %+v", err)
	}
	if result.Error != "" {
		return nil, fmt.Errorf("get playlists fault: %s", result.Error)
	}
	return result.Result.Tracks, nil
}

/*
 * Tracks
 */

func (c *Client) LosslessInfo(ctx context.Context, trackId string) (*ent.TrackLosslessInfo, error) {
	var (
		result ent.TrackLosslessInfoResp
		ts     = strconv.FormatInt(time.Now().Unix(), 10)
	)

	params := url.Values{
		"ts":         {ts},
		"trackId":    {trackId},
		"quality":    {"lossless"},
		"codecs":     {"flac,aac,he-aac,mp3"},
		"transports": {"raw"}}

	if sign, err := signData([]byte(fmt.Sprintf("%s%slosslessflacaache-aacmp3raw", ts, trackId))); err != nil {
		return nil, err
	} else {
		params.Add("sign", sign)
	}

	if err := (*c).callAPI(ctx, http.MethodGet, "get-file-info", params, nil, &result); err != nil {
		return nil, fmt.Errorf("lossless info fault: %+v", err)
	}
	if result.Error != "" {
		return nil, fmt.Errorf("lossless inf fault: %s", result.Error)
	}
	return &result.Result.Info, nil
}

func (c *Client) DownloadTrack(ctx context.Context, t *ent.Track) error {
	var (
		exist     bool
		sb        strings.Builder
		delimiter = []byte{44, 32}
		aCount    = len((*t).Artists)
	)

	for i, a := range (*t).Artists {
		sb.WriteString(a.Name)
		if i < aCount-1 {
			sb.Write(delimiter)
		}
	}
	artists := sb.String()

	// cover image
	dirName := fmt.Sprintf("%s/images/%s", (*c).cacheDir, artists)
	fName := fmt.Sprintf("%s.png", (*t).Title)
	if exist, err := checkFile(dirName, fName); err != nil {
		return err
	} else if !exist {
		if err = (*c).downloadImage((*t).OgImage, fmt.Sprintf("%s/%s", dirName, fName)); err != nil {
			return err
		}
	}

	// audio file
	if info, err := (*c).LosslessInfo(ctx, (*t).Id); err != nil {
		return err
	} else {
		dirName = fmt.Sprintf("%s/tracks/%s", (*c).cacheDir, artists)
		fName = fmt.Sprintf("%s.%s", (*t).Title, (*info).Codec)
		if exist, err = checkFile(dirName, fName); err != nil {
			return err
		} else if !exist {
			if err = (*c).download((*info).Url, fmt.Sprintf("%s/%s", dirName, fName), nil); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) downloadImage(uri, fName string) error {
	uri = fmt.Sprintf("https://%s", strings.Replace(uri, "%%", "200x200", 1))
	return (*c).download(uri, fName, nil)
}

func (c *Client) getDirectLink(infoUrl string) (string, error) {
	// TODO: check flac download
	var (
		data interface{}
		buf  = new(bytes.Buffer)
	)

	if err := (*c).download(infoUrl, "", buf); err != nil {
		return "", err
	}

	if err := xml.Unmarshal(buf.Bytes(), &data); err != nil {
		return "", err
	}

	log.Printf("XML: %+v", data)
	// The URL available for one minute after receiving the download information.
	// Otherwise HTTP 410 error will be received.

	return "", nil
}

func checkFile(dirName, fileName string) (bool, error) {
	if err := os.MkdirAll(dirName, 0750); err != nil {
		return false, err
	}
	if _, err := os.Stat(fmt.Sprintf("%s/%s", dirName, fileName)); errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func signData(data []byte) (string, error) {
	mac := hmac.New(sha256.New, AndroidKey)
	if _, err := mac.Write(data); err != nil {
		return "", err
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, mac.Sum(nil))
	dst = bytes.TrimRight(dst, "\x00")
	return strings.TrimSuffix(string(dst), "="), nil
}

//
// func (c *Client) PermissionAlerts() (*ent.PermissionAlerts, error) {
// 	var result entities.PermissionAlertsResult
// 	url, err := url.JoinPath(c.baseUrl, "/permission-alerts")
// 	if err != nil {
// 		return result.Result, err
// 	}
//
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		return result.Result, err
// 	}
// 	for k, v := range headers {
// 		req.Header.Set(k, v)
// 	}
// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return result.Result, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return result.Result, fmt.Errorf("yandex music api returned wrong code:" + strconv.Itoa(resp.StatusCode))
// 	}
// 	err = result.ParseFromReader(resp.Body)
// 	return result.Result, err
// }
//
// func (c *Client) AccountExperiments() (*ent.AccountExperiments, error) {
// 	var result entities.AccountExperimentsResult
// 	url, err := url.JoinPath(c.baseUrl, "/account/experiments")
// 	if err != nil {
// 		return result.Result, err
// 	}
//
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		return result.Result, err
// 	}
// 	for k, v := range headers {
// 		req.Header.Set(k, v)
// 	}
// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return result.Result, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return result.Result, fmt.Errorf("yandex music api returned wrong code:" + strconv.Itoa(resp.StatusCode))
// 	}
// 	err = result.ParseFromReader(resp.Body)
// 	return result.Result, err
// }
//
// func (c *Client) ConsumePromocode(promocode, language string) (*ent.PromocodeStatus, error) {
// 	var result entities.PromocodeStatusResult
// 	url, err := url.JoinPath(c.baseUrl, "/account/consume-promo-code")
// 	if err != nil {
// 		return result.Result, err
// 	}
// 	data, err := json.Marshal(map[string]string{"code": promocode, "language": language})
// 	if err != nil {
// 		return result.Result, err
// 	}
//
// 	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
// 	if err != nil {
// 		return result.Result, err
// 	}
// 	for k, v := range headers {
// 		req.Header.Set(k, v)
// 	}
// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return result.Result, err
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return result.Result, fmt.Errorf("yandex music api returned wrong code:" + strconv.Itoa(resp.StatusCode))
// 	}
// 	err = result.ParseFromReader(resp.Body)
// 	return result.Result, err
// }
