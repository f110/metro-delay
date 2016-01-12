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

type SlackTeam struct {
	Name        string `json:"name"`
	AccessToken string `json:"access_token"`
}

type WatcherConf struct {
	Channels         []ChannelConf
	SlackTeams       []SlackTeam `json:"slack_teams"`
	MetroAccessToken string      `json:"metro_access_token"`
	SlackTeamMap     map[string]SlackTeam
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

	conf.SlackTeamMap = make(map[string]SlackTeam)
	for _, v := range conf.SlackTeams {
		conf.SlackTeamMap[v.Name] = v
	}

	return &conf, nil
}

func (watcherConf *WatcherConf) GetAccessToken(team string) string {
	if v, ok := watcherConf.SlackTeamMap[team]; ok {
		return v.AccessToken
	}

	return ""
}
