package twitter

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/dghubble/oauth1"
	"github.com/google/go-querystring/query"
)

const twitterAPI = "https://api.twitter.com/1.1"

type Client struct {
	APIClient APIClient

	TimelineService *TimelineService
	StatusService   *StatusService
}

func NewClient(config *Config) *Client {
	apiClient := newAPIClient(config)

	return &Client{
		APIClient: apiClient,

		TimelineService: newTimelineService(apiClient),
		StatusService:   newStatusService(apiClient),
	}
}

type APIClient interface {
	Get(uri string, values interface{}) ([]byte, int, error)
	Post(uri string, values interface{}) ([]byte, int, error)
}

type apiClient struct {
	HttpClient *http.Client
}

func newAPIClient(config *Config) *apiClient {
	oauthConfig := oauth1.NewConfig(config.ConsumerApiKey, config.ConsumerApiSecret)
	token := oauth1.NewToken(config.AccessToken, config.AccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	return &apiClient{
		HttpClient: httpClient,
	}
}

func (apiClient *apiClient) Get(uri string, values interface{}) ([]byte, int, error) {
	q, err := query.Values(values)
	if err != nil {
		return nil, 0, err
	}

	u, err := url.Parse(twitterAPI)
	if err != nil {
		return nil, 0, err
	}

	u.Path = path.Join(u.Path, uri)
	u.RawQuery = q.Encode()

	resp, err := apiClient.HttpClient.Get(u.String())
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}

func (apiClient *apiClient) Post(uri string, values interface{}) ([]byte, int, error) {
	q, err := query.Values(values)
	if err != nil {
		return nil, 0, err
	}

	u, err := url.Parse(twitterAPI)
	if err != nil {
		return nil, 0, err
	}

	u.Path = path.Join(u.Path, uri)

	resp, err := apiClient.HttpClient.Post(u.String(), "application/x-www-form-urlencoded", strings.NewReader(q.Encode()))
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return body, resp.StatusCode, nil
}
