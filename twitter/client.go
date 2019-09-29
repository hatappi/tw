package twitter

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/dghubble/oauth1"
	"github.com/google/go-querystring/query"
)

const twitterAPI = "https://api.twitter.com/1.1"

type Client struct {
	APIClient APIClient

	TimelineService *TimelineService
}

func NewClient(config *Config) *Client {
	apiClient := newAPIClient(config)

	return &Client{
		APIClient: apiClient,

		TimelineService: newTimelineService(apiClient),
	}
}

type APIClient interface {
	Get(uri string, values interface{}) ([]byte, int, error)
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
