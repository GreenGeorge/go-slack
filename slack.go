package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const slackBaseURL = "https://slack.com/api"

type Slack struct {
	clientID     string
	clientSecret string
	accessToken  string
	client       *http.Client
}

// New assembles a new Slack instance loaded with credentials
func New(accessToken string, client *http.Client) Slack {
	return Slack{
		accessToken: accessToken,
		client:      client,
	}
}

// PostMessage posts a formatted message to a Slack workspace
func (s *Slack) PostMessage(message Message) (*http.Response, error) {
	json, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	payloadBuffer := bytes.NewBuffer(json)
	endpoint := "/chat.postMessage"
	targetURL := fmt.Sprintf("%s%s", slackBaseURL, endpoint)
	req, err := http.NewRequest(http.MethodPost, targetURL, payloadBuffer)
	if err != nil {
		return nil, err
	}
	AuthHeader := fmt.Sprintf("Bearer %s", s.accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", AuthHeader)
	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
