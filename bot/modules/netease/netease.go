package netease

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	baseURL = "https://163.lpddr5.cn"
)

type Payload = map[string]interface{}

func makeRequest(endpoint string, payload Payload) ([]byte, error) {
	res, body, errs := gorequest.New().
		Post(fmt.Sprintf("%s/%s", baseURL, endpoint)).
		Send(payload).
		EndBytes()
	if errs != nil {
		return nil, fmt.Errorf("makeRequest: %s: %v", endpoint, errs)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("makeRequest: %s: %d %s", endpoint, res.StatusCode, res.Status)
	}
	return body, nil
}

func Search(keyword string) (*SearchResult, error) {
	query := fmt.Sprintf("search?timestamp=%d", time.Now().Unix())
	bytes, err := makeRequest(query, Payload{
		"keywords": keyword,
	})
	if err != nil {
		return nil, err
	}
	var res SearchResult
	_ = json.Unmarshal(bytes, &res)
	return &res, nil
}

func GetSongDetail(id int64) (*SongDetail, error) {
	bytes, err := makeRequest("song/url", Payload{
		"id": id,
	})
	if err != nil {
		return nil, err
	}
	var res SongDetail
	_ = json.Unmarshal(bytes, &res)
	if res.Url == "" {
		return nil, fmt.Errorf("音乐无版权、或者不存在")
	}
	return &res, nil
}
