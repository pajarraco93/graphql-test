package lastfm

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pajarraco93/graphql-test/pkg/library/domain/entities"
)

type searchParams map[string]string

type lastFMAPI struct {
	key string
	url string
}

const UserAgent = "Dataquest"

func ErrUnexpectedStatusCode(status string) error {
	return fmt.Errorf("lastFM: unexpected response %s", status)
}

func NewLastFMAPI(apiKey, url string) *lastFMAPI {
	return &lastFMAPI{
		key: apiKey,
		url: url,
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

	req.Header.Set("user-agent", UserAgent)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return "", ErrUnexpectedStatusCode(res.Status)
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
