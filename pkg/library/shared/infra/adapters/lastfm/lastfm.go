package lastfm

import (
	"io/ioutil"
	"net/http"

	"github.com/pajarraco93/graphql-test/pkg/library/domain/entities"
)

var USER_AGENT = "Dataquest"

type searchParams map[string]string

type lastFMAPI struct {
	key string
	url string
}

func NewLastFMAPI(apiKey string) *lastFMAPI {
	return &lastFMAPI{
		key: apiKey,
		url: "http://ws.audioscrobbler.com/2.0/",
	}
}

func (api *lastFMAPI) getCall(method string, extraParams searchParams) (string, error) {
	req, err := http.NewRequest("GET", api.url, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("api_key", api.key)
	q.Add("format", "json")
	q.Add("method", method)

	for key, value := range extraParams {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	req.Header.Set("user-agent", USER_AGENT)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)

	return string(body), nil
}

func (api *lastFMAPI) GetGroupInfo(group entities.Group) (string, error) {
	params := searchParams{
		"artist":      group.Name,
		"autocorrect": "1",
	}
	return api.getCall("artist.getInfo", params)
}

func (api *lastFMAPI) GetAlbumInfo(album entities.Album) (string, error) {
	params := searchParams{
		"artist":      album.Name,
		"autocorrect": "1",
	}
	return api.getCall("album.getInfo", params)
}

func (api *lastFMAPI) GetSongInfo(song entities.Song) (string, error) {
	params := searchParams{
		"artist":      song.Name,
		"autocorrect": "1",
	}
	return api.getCall("track.getInfo", params)
}
