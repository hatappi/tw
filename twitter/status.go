package twitter

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}

type Place struct {
	Attributes  map[string]string `json:"attributes"`
	BoundingBox *BoundingBox      `json:"bounding_box"`
	Country     string            `json:"country"`
	CountryCode string            `json:"country_code"`
	FullName    string            `json:"full_name"`
	Geometry    *BoundingBox      `json:"geometry"`
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	PlaceType   string            `json:"place_type"`
	Polylines   []string          `json:"polylines"`
	URL         string            `json:"url"`
}

type BoundingBox struct {
	Coordinates [][][2]float64 `json:"coordinates"`
	Type        string         `json:"type"`
}

type Tweet struct {
	CreatedAt           string      `json:"created_at"`
	ID                  int         `json:"id"`
	Text                string      `json:"text"`
	Source              string      `json:"source"`
	Truncated           bool        `json:"truncated"`
	InReplyToStatusID   int         `json:"in_reply_to_status_id"`
	InReplyToUserID     int         `json:"in_reply_to_user_id"`
	InReplyToScreenName string      `json:"in_reply_to_screen_name"`
	User                User        `json:"user"`
	Coordinates         Coordinates `json:"coordinates"`
	Place               Place       `json:"place"`
	QuotedStatusID      int         `json:"quoted_status_id"`
	IsQuoteStatus       bool        `json:"is_quote_status"`
	QuotedStatus        *Tweet      `json:"quoted_status"`
	RetweetedStatus     *Tweet      `json:"retweeted_status"`
	QuoteCount          int         `json:"quote_count"`
	ReplyCount          int         `json:"reply_count"`
	RetweetCount        int         `json:"retweet_count"`
	FavoriteCount       int         `json:"favorite_count"`
	Favorited           bool        `json:"favorited"`
	Retweeted           bool        `json:"retweeted"`
	PossiblySensitive   bool        `json:"possibly_sensitive"`
	Lang                string      `json:"lang"`
}

func (t *Tweet) CreatedAtTime() (time.Time, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return time.Time{}, err
	}

	c, err := time.Parse(time.RubyDate, t.CreatedAt)
	if err != nil {
		return time.Time{}, err
	}

	return c.In(jst), nil
}

type StatusService struct {
	apiClient APIClient
}

func newStatusService(apiClient APIClient) *StatusService {
	return &StatusService{
		apiClient: apiClient,
	}
}

// https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/post-statuses-update
type UpdateStatusParams struct {
	Status                    string  `url:"status"`
	InReplyToStatusID         int     `url:"in_reply_to_status_id,omitempty"`
	AutoPopulateReplyMetadata bool    `url:"auto_populate_reply_metadata,omitempty"`
	ExcludeReplyUserIDs       []int   `url:"exclude_reply_user_ids,omitempty"`
	AttachmentUrl             string  `url:"attachment_url,omitempty"`
	MediaIDs                  []int   `url:"media_ids,omitempty"`
	PossiblySensitive         bool    `url:"possibly_sensitive,omitempty"`
	Lat                       float32 `url:"lat,omitempty"`
	Long                      float32 `url:"long,omitempty"`
	PlaceID                   string  `url:"place_id,omitempty"`
	DisplayCoordinates        bool    `url:"display_coordinates,omitempty"`
	TrimUser                  bool    `url:"trim_user,omitempty"`
	EnableDmcommands          bool    `url:"enable_dmcommands,omitempty"`
	FailDmcommands            bool    `url:"fail_dmcommands,omitempty"`
	CardUri                   string  `url:"card_uri,omitempty"`
}

func (serv *StatusService) UpdateStatus(params *UpdateStatusParams) (*Tweet, error) {
	if params.Status == "" {
		return nil, fmt.Errorf("please set Status")
	}

	body, _, err := serv.apiClient.Post("statuses/update.json", params)
	if err != nil {
		return nil, err
	}

	var tweet *Tweet
	err = json.Unmarshal(body, &tweet)
	if err != nil {
		return nil, err
	}

	return tweet, nil
}
