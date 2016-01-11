package main

import (
	"encoding/json"
	"io/ioutil"
)

type ChannelConf struct {
	Railway    string
	Channel    []string
	AuthorIcon string `json:"author_icon"`
	ThumbUrl   string `json:"thumb_url"`
	IconUrl    string `json:"icon_url"`
}

type WatcherConf struct {
	Channels         []ChannelConf
	MetroAccessToken string `json:"metro_access_token"`
	SlackAccessToken string `json:"slack_access_token"`
}

func NewConf(filePath string) (*WatcherConf, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var conf WatcherConf
	err = json.Unmarshal(buf, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
