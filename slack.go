package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Slack struct {
	clientID     string
	clientSecret string
	accessToken  string
	client       *http.Client
}

// TODO Double check and complete message schema

type Message struct {
	Channel         string       `json:"channel,omitempty"`
	Text            string       `json:"text,omitempty"`
	UserName        string       `json:"username,omitempty"`
	Attachments     []Attachment `json:"attachments,omitempty"`
	ThreadTS        string       `json:"thread_ts,omitempty"`
	ResponseType    string       `json:"response_type,omitempty"`
	ReplaceOriginal bool         `json:"replace_original,omitempty"`
	DeleteOriginal  bool         `json:"delete_original,omitempty"`
	Mrkdwn          bool         `json:"mrkdwn,omitempty"`
}

type Attachment struct {
	Fallback       string   `json:"fallback,omitempty"`
	CallbackID     string   `json:"callback_id,omitempty"`
	Color          string   `json:"color,omitempty"`
	Actions        []Action `json:"actions,omitempty"`
	AttachmentType string   `json:"attachment_type,omitempty"`
	PreText        string   `json:"pretext,omitempty"`
	AuthorName     string   `json:"author_name,omitempty"`
	AuthorLink     string   `json:"author_link,omitempty"`
	AuthorIcon     string   `json:"author_icon,omitempty"`
	Title          string   `json:"title,omitempty"`
	TitleLink      string   `json:"title_link,omitempty"`
	Text           string   `json:"text,omitempty"`
	MrkdwnIn       []string `json:"mrkdwn_in,omitempty"`
	Fields         []Field  `json:"fields,omitempty"`
	ImageURL       string   `json:"image_url,omitempty"`
	ThumbURL       string   `json:"thumb_url,omitempty"`
	Footer         string   `json:"footer,omitempty"`
	FooterIcon     string   `json:"footer_icon,omitempty"`
	TS             int      `json:"ts,omitempty"`
}

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

type Action struct {
	Name           string `json:"name,omitempty"`
	Text           string `json:"text,omitempty"`
	Type           string `json:"type,omitempty"`
	Value          string `json:"value,omitempty"`
	Confirm        `json:"confirm,omitempty"`
	Style          string        `json:"style,omitempty"`
	Options        []Option      `json:"options,omitempty"`
	OptionGroups   []OptionGroup `json:"option_groups,omitempty"`
	DataSource     string        `json:"data_source,omitempty"`
	SelectdOptions []Option      `json:"selected_options,omitempty"`
	MinQueryLength int           `json:"min_query_length,omitempty"`
	URL            string        `json:"url,omitempty"`
}

type Confirm struct {
	Title       string `json:"title,omitempty"`
	Text        string `json:"text,omitempty"`
	OKText      string `json:"ok_text,omitempty"`
	DismissText string `json:"dismiss_text,omitempty"`
}

type OptionGroup struct {
	Text string `json:"text,omitempty"`
	Option
}

type Option struct {
	Text        string `json:"text,omitempty"`
	Value       string `json:"text,omitempty"`
	Description string `json:"text,omitempty"`
}

// New assembles a new Slack instance loaded with credentials
func New(accessToken string, client *http.Client) Slack {
	return Slack{
		accessToken: accessToken,
		client:      client,
	}
}

// PostMessage posts a formatted message to a Slack workspace
func (s *Slack) PostMessage(webhookURL string, message Message) (*http.Response, error) {
	json, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	payloadBuffer := bytes.NewBuffer(json)
	req, err := http.NewRequest(http.MethodPost, webhookURL, payloadBuffer)
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
