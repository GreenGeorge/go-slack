package slack

// TODO Double check and complete message schema

type Message struct {
	Channel         string       `json:"channel,omitempty"`
	Text            string       `json:"text,omitempty"`
	UserName        string       `json:"username,omitempty"`
	Attachments     []Attachment `json:"attachments,omitempty"`
	IconEmoji       string       `json:"icon_emoji,omitempty"`
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
	Name           string        `json:"name,omitempty"`
	Text           string        `json:"text,omitempty"`
	Type           string        `json:"type,omitempty"`
	Value          string        `json:"value,omitempty"`
	Confirm        Confirm       `json:"confirm,omitempty"`
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

// TODO flesh out message constructor

// NewBaseMessage creates a new base message pre-loaded with sender info and message destination
func NewBaseMessage(username, channel, emoji string) Message {
	return Message{
		UserName:  username,
		Channel:   channel,
		IconEmoji: emoji,
	}
}

// AddAttachment adds an attachment onto a message
func (m *Message) AddAttachment(attachments ...Attachment) *Message {
	for _, attachment := range attachments {
		m.Attachments = append(m.Attachments, attachment)
	}
	return m
}

// AddAction adds an action onto an attachment
func (a *Attachment) AddAction(actions ...Action) *Attachment {
	for _, action := range actions {
		a.Actions = append(a.Actions, action)
	}
	return a
}

// AddField adds a field onto an attachment
func (a *Attachment) AddField(fields ...Field) *Attachment {
	for _, field := range fields {
		a.Fields = append(a.Fields, field)
	}
	return a
}
