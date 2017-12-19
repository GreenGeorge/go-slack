package slack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBaseMessage(t *testing.T) {
	tests := map[string]struct {
		username string
		channel  string
		emoji    string
	}{
		"Valid strings passed": {
			username: "George",
			channel:  "123456",
			emoji:    ":rocket:",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			m := NewBaseMessage(test.username, test.channel, test.emoji)
			assert.Equal(t, test.username, m.UserName)
			assert.Equal(t, test.channel, m.Channel)
			assert.Equal(t, test.emoji, m.IconEmoji)
		})
	}
}

func TestAddAttachments(t *testing.T) {
	tests := map[string]struct {
		attachments []Attachment
	}{
		"Adds a single attachment": {
			attachments: []Attachment{
				Attachment{},
			},
		},
		"Adds more than one attachment": {
			attachments: []Attachment{
				Attachment{},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			m := Message{}
			m.AddAttachment(test.attachments...)
			expected := len(test.attachments)
			actual := len(m.Attachments)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestAddActions(t *testing.T) {
	tests := map[string]struct {
		actions []Action
	}{
		"Adds a single action": {
			actions: []Action{
				Action{},
			},
		},
		"Adds more than one action": {
			actions: []Action{
				Action{},
				Action{},
				Action{},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := Attachment{}
			a.AddAction(test.actions...)
			expected := len(test.actions)
			actual := len(a.Actions)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestAddField(t *testing.T) {
	tests := map[string]struct {
		fields []Field
	}{
		"Adds a single field": {
			fields: []Field{
				Field{},
			},
		},
		"Adds more than one field": {
			fields: []Field{
				Field{},
				Field{},
				Field{},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			a := Attachment{}
			a.AddField(test.fields...)
			expected := len(test.fields)
			actual := len(a.Fields)
			assert.Equal(t, expected, actual)
		})
	}

}
