package twitter

import (
	"encoding/json"
)

type TimelineService struct {
	apiClient APIClient
}

func newTimelineService(apiClient APIClient) *TimelineService {
	return &TimelineService{
		apiClient: apiClient,
	}
}

// https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-home_timeline#parameters
type HomeTimelineParams struct {
	Count           int  `url:"count,omitempty"`
	SinceID         int  `url:"since_id,omitempty"`
	MaxID           int  `url:"max_id,omitempty"`
	TrimUser        bool `url:"trim_user,omitempty"`
	ExcludeReplies  bool `url:"exclude_replies,omitempty"`
	IncludeEntities bool `url:"include_entities,omitempty"`
}

func (serv *TimelineService) GetHomeTimeline(params *HomeTimelineParams) ([]*Tweet, error) {
	body, _, err := serv.apiClient.Get("statuses/home_timeline.json", params)
	if err != nil {
		return nil, err
	}

	var tweets []*Tweet
	err = json.Unmarshal(body, &tweets)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
