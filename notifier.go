package main

import (
    "errors"
    "encoding/json"
    "net/http"
    "net/url"
)

type Attachment struct {
	Color string `json:"color"`
	Text string `json:"text"`
	Fallback string `json:"fallback"`
	AuthorName string `json:"author_name"`
	AuthorIcon string `json:"author_icon"`
	ThumbUrl string `json:"thumb_url"`
}

type PostMessageResponse struct {
	Ok bool
    Error string
}

type SlackNotifier struct {
    AccessToken string
    Channels []ChannelConf
}

func NewSlackNotifier(conf *WatcherConf) (*SlackNotifier, error) {
    return &SlackNotifier{AccessToken: conf.SlackAccessToken, Channels: conf.Channels}, nil
}

func (slackNotifier *SlackNotifier) Notify(railway string, text string) error {
	attachment := &Attachment{
		Color: "good",
		Text: text,
		Fallback: text,
		AuthorName: RailwayToName[railway],
		// AuthorIcon: "http://go-imghr.ds-12.com/f.jpg",
		// ThumbUrl: "http://go-imghr.ds-12.com/f.jpg",
	}
	attachments := make([]Attachment, 1)
	attachments[0] = *attachment
	buf, err := json.Marshal(attachments)
	if err != nil {
		return err
	}

    channel := ""
    for _, v := range(slackNotifier.Channels) {
        if v.Railway == railway {
            channel = v.Channel
            break
        }
    }
    if channel == "" {
        return errors.New("Could not find channel")
    }

	v := url.Values{}
	v.Set("token", slackNotifier.AccessToken)
	v.Set("channel", channel)
	v.Set("text", "")
	v.Set("username", "metrobot")
	v.Set("attachments", string(buf))

	response, _ := http.PostForm("https://slack.com/api/chat.postMessage", v)
	dec := json.NewDecoder(response.Body)
	var data PostMessageResponse
	dec.Decode(&data)

	if data.Ok != true {
        return errors.New("Failed postMessage " + data.Error)
	}

    return nil
}
