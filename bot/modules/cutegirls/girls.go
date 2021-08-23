package cutegirls

import (
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
)

const (
	girlsApi  = "https://api.dongmanxingkong.com/suijitupian/acg/1080p/index.php?return=json"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"
)

type CuteGirl struct {
	//200
	Code string `json:"code"`
	//example "https://tvax4.sinaimg.cn/large/006ZFECEgy1fr1x6fq53hj31hc0u0h87.jpg",
	Url string `json:"imgurl"`
	//"1920",
	Width string `json:"width"`
	//"1080"
	Height string `json:"height"`
}

type Image struct {
	Url  string
	Body []byte
}

func WantCuteGirl() (*Image, error) {
	girl := CuteGirl{}
	_, _, errs := gorequest.New().
		Get(girlsApi).
		Set("User-Agent", userAgent).
		EndStruct(&girl)
	if errs != nil {
		return nil, fmt.Errorf("girls api: %v", errs)
	}
	if girl.Code != "200" {
		return nil, fmt.Errorf("girls api code: %v", girl.Code)
	}
	logger.Infof("got girlapi result: %v", girl.Url)
	startTime := time.Now()
	resp, bytes, errs := gorequest.New().
		Get(girl.Url).
		Set("User-Agent", userAgent).
		EndBytes()
	spanTime := time.Since(startTime)
	if errs != nil {
		return nil, fmt.Errorf("girls bytes: %v", errs)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("girls bytes status: %v", resp.StatusCode)
	}
	logger.Infof("fetch image consumption: %s", spanTime)
	return &Image{Body: bytes, Url: girl.Url}, nil
}
