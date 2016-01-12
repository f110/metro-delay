package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

//good, warning or danger, or color code
var StatusColormap = map[string]string{
	StatusSuspended:         "danger",
	StatusShuttle:           "danger",
	StatusTimetableDisarray: "warning",
	StatusDelay:             "warning",
	StatusPartiallyDelay:    "warning",
	StatusAbortDirect:       "warning",
	StatusRestartDirect:     "good",
	StatusCancel:            "danger",
}

type Attachment struct {
	Color      string `json:"color"`
	Text       string `json:"text"`
	Fallback   string `json:"fallback"`
	AuthorName string `json:"author_name"`
	AuthorIcon string `json:"author_icon"`
	ThumbUrl   string `json:"thumb_url"`
}

type PostMessageResponse struct {
	Ok    bool
	Error string
}

type SlackNotifier struct {
	Conf *WatcherConf
}

func NewSlackNotifier(conf *WatcherConf) (*SlackNotifier, error) {
	return &SlackNotifier{Conf: conf}, nil
}

func (slackNotifier *SlackNotifier) Notify(railway string, text string, status string) error {
	attachment := &Attachment{
		Color:      "good",
		Text:       text,
		Fallback:   text,
		AuthorName: RailwayToName[railway],
	}

	if status != "" {
		if color, ok := StatusColormap[status]; ok {
			attachment.Color = color
		}
	}

	var channel []string
	var iconUrl string
	for _, v := range slackNotifier.Conf.Channels {
		if v.Railway == railway {
			if v.AuthorIcon != "" {
				attachment.AuthorIcon = v.AuthorIcon
			}
			if v.ThumbUrl != "" {
				attachment.ThumbUrl = v.ThumbUrl
			}
			if v.IconUrl != "" {
				iconUrl = v.IconUrl
			}
			channel = v.Channel
			break
		}
	}
	if len(channel) == 0 {
		return errors.New("Could not find channel")
	}

	attachments := make([]Attachment, 1)
	attachments[0] = *attachment
	buf, err := json.Marshal(attachments)
	if err != nil {
		return err
	}
	body := url.Values{}
	body.Set("text", "")
	body.Set("username", "metrobot")
	if iconUrl != "" {
		body.Set("icon_url", iconUrl)
	}
	body.Set("attachments", string(buf))

	for _, v := range channel {
		teamAndChannel := strings.Split(v, ":")
		if len(teamAndChannel) != 2 {
			continue
		}
		body.Set("token", slackNotifier.Conf.GetAccessToken(teamAndChannel[0]))
		body.Set("channel", teamAndChannel[1])
		response, _ := http.PostForm("https://slack.com/api/chat.postMessage", body)
		dec := json.NewDecoder(response.Body)
		var data PostMessageResponse
		dec.Decode(&data)

		if data.Ok != true {
			return errors.New("Failed postMessage " + data.Error)
		}
	}

	return nil
}
