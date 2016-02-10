package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type MetroWatcher struct {
	form           url.Values
	lastStatusTime map[string]TrainInformationResponse
}

type TrainInformationResponse struct {
	Context      string    `json:"@context"`
	Id           string    `json:"@id"`
	Type         string    `json:"@type"`
	Date         time.Time `json:"dc.date"`
	Valid        time.Time `json:"dct:valid"`
	Operator     string    `json:"odpt:operator"`
	TimeOfOrigin time.Time `json:"odpt:timeOfOrigin"`
	Railway      string    `json:"odpt:railway"`
	Status       string    `json:"odpt:trainInformationStatus"`
	Text         string    `json:"odpt:trainInformationText"`
}

func NewMetroWatcher(conf *WatcherConf) (*MetroWatcher, error) {
	form := url.Values{}
	form.Add("rdf:type", TypeTrainInfomation)
	form.Add("acl:consumerKey", conf.MetroAccessToken)

	lastStatusTime := make(map[string]TrainInformationResponse)
	return &MetroWatcher{form: form, lastStatusTime: lastStatusTime}, nil
}

func (watcher *MetroWatcher) Get(retry int) ([]TrainInformationResponse, error) {
	if retry < 1 {
		return nil, errors.New("give up")
	}

	res, err := http.Get("https://api.tokyometroapp.jp/api/v2/datapoints?" + watcher.form.Encode())
	if retry < 1 && err != nil {
		return nil, err
	}
	// リトライできる時はエラーを返さずにリトライする
	if err != nil {
		time.Sleep(1 * time.Minute)
		return watcher.Get(retry - 1)
	}
	if res.StatusCode != 200 {
		time.Sleep(10 * time.Second)
		return watcher.Get(retry - 1)
	}

	buf, err := ioutil.ReadAll(res.Body)
	jsonBuf := make([]TrainInformationResponse, 0)
	err = json.Unmarshal(buf, &jsonBuf)
	if err != nil {
		log.Print(err)
		return jsonBuf, nil
	}

	return jsonBuf, nil
}

func (watcher *MetroWatcher) Start(notifier *SlackNotifier) error {
	for {
		current, err := watcher.Get(3)
		if err != nil {
			return err
		}

		// FIXME: v.Railwayのvalidation
		for _, v := range current {
			if _, ok := watcher.lastStatusTime[v.Railway]; ok == false {
				watcher.lastStatusTime[v.Railway] = v
				continue
			}

			if watcher.shouldNotify(v) {
				log.Print(watcher.lastStatusTime)
				notifier.Notify(v.Railway, v.Text, v.Status)
				watcher.lastStatusTime[v.Railway] = v
			}
		}

		time.Sleep(5 * time.Minute)
	}
}

func (watcher *MetroWatcher) shouldNotify(v TrainInformationResponse) bool {
	// ステータスに何らかの値が入ってて（=異常）TimeOfOriginが違ったらtrue
	if v.Status != "" && watcher.lastStatusTime[v.Railway].TimeOfOrigin.Unix() != v.TimeOfOrigin.Unix() {
		return true
	}

	// ステータスありからステータスなしに変わった場合はtrue
	if watcher.lastStatusTime[v.Railway].Status != "" && v.Status == "" {
		return true
	}

	return false
}
