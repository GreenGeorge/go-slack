Go-Slack
========
[![CircleCI](https://circleci.com/gh/GreenGeorge/go-slack/tree/master.svg?style=shield)](https://circleci.com/gh/GreenGeorge/go-slack/tree/master)
[![GoDoc](https://godoc.org/github.com/greengeorge/go-slack?status.svg)](http://godoc.org/github.com/greengeorge/go-slack)

Really simple Go library for programatically creating and delivering Slack messages. Just BYO [access token][2] and `http.Client`. Great for creating `slack` reports with reusable components.

## Usage

Install it `$ go get github.com/greengeorge/go-slack`

```go
...
import "github.com/GreenGeorge/go-slack"

// Bring your own Client or just pass nil
client := http.Client{Timeout: time.Second * 10}

// Instantiate go-slack with your access token)
sl := slack.New("xoxp-XXXXXXXXXXXXXX-XXXXXXXXXXXXXX-XXXXXXXXXXXXXX", client)

// Prepare attachments
attachmentFoo := slack.Attachment{
  Title:          "Foo",
  Text:           "You've got Foo",
  Color:          "#f5b000",
  AttachmentType: "default",
  Fallback:       "You've got Foo",
}
attachmentBar := slack.Attachment{
  Title:          "Bar",
  Text:           "You've got Bar",
  Color:          "#00e6f5",
  AttachmentType: "default",
  Fallback:     "You've got Bar",
}

// Prepare reusable actions
actionBaz := slack.Action{
  Name:   "Baz",
  Text:   "Send out the Baz",
  Type:   "button",
  Value:  "baz",
  Style:  "primary",
}
actionFiz := slack.Action{
  Name:   "Fiz",
  Text:   "Send out the Fiz",
  Type:   "button",
  Value:  "fiz",
  Style:  "danger",
}

// Mix and match them!
attachmentFoo.AddAction(actionBaz, actionFiz)
attachmentBar.AddAction(actionFiz)

// Setup base messages
messageA := slack.NewMessage("FooBar", "CHANNELIDA", ":skull:")
messageB := slack.NewMessage("BazFiz", "CHANNELIDB",":skull:")

// Assemble the final message
// Add a text message maybe
messageA.Text = "Kunci kesuksesan adalah key of success"

// Add the attachments to the message
messageA.AddAttachment(attachmentFoo, attachmentBar)
messageB.AddAttachment(attachmentBar, attachmentFoo)

// Send them out!
_, err := sl.PostMessage(messageA)
_, err := sl.PostMessage(messageB)

...
```

*et voilà*

![Example result][example]

## API Coverage
Definitely doesn't cover the whole Slack API. Just the parts I need for my use cases (easy message creation and sending by Slack's [Web API][1]). Contributions are most welcome.

Please note the message struct schema might also be incomplete. I've yet to find one list that covers everything.

[1]:https://www.google.com
[2]:https://api.slack.com/apps
[example]:example/example-result.jpg
