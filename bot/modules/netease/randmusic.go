package netease

import (
	"fmt"
	"math/rand"

	"github.com/parnurzeal/gorequest"
)

type RandomSong struct {
	Code int `json:"code"`
	Data struct {
		Name   string `json:"name"`
		Url    string `json:"url"`
		PicUrl string `json:"picurl"`
		Artist string `json:"artistsname"`
	} `json:"data"`
}

func GetRandomSong() (*RandomSong, error) {
	categories := []string{"热歌榜", "新歌榜", "飙升榜", "抖音榜", "电音榜"}
	url := fmt.Sprintf(
		"https://api.uomg.com/api/rand.music?sort=%s&format=json",
		categories[rand.Intn(len(categories))])
	var res RandomSong
	resp, _, errs := gorequest.New().Get(url).EndStruct(&res)
	if errs != nil {
		return nil, fmt.Errorf("get random song: %v", errs)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("get random song: %d %v", resp.StatusCode, resp.Status)
	}
	logger.Infof("get random song: %s %s", res.Data.Name, res.Data.Url)
	return &res, nil
}
