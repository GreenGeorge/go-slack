package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	slackBaseURL        = "https://slack.com/api"
	endpointPostMessage = "/chat.postMessage"
)

type Client struct {
	accessToken string
	httpClient  *http.Client
}

func (c Client) AccessToken() string {
	return c.accessToken
}

func (c Client) HttpClient() *http.Client {
	return c.httpClient
}

// New assembles a new Slack instance loaded with credentials
func New(accessToken string, httpClient *http.Client) Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return Client{
		accessToken: accessToken,
		httpClient:  httpClient,
	}
}

// PostMessage posts a formatted message to a Slack workspace
func (s *Client) PostMessage(message Message) (*http.Response, error) {
	json, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	payloadBuffer := bytes.NewBuffer(json)
	targetURL := fmt.Sprintf("%s%s", slackBaseURL, endpointPostMessage)
	req, err := http.NewRequest(http.MethodPost, targetURL, payloadBuffer)
	if err != nil {
		return nil, err
	}
	AuthHeader := fmt.Sprintf("Bearer %s", s.accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", AuthHeader)
	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
