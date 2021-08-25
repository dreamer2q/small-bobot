package utils

import (
	"fmt"
	"log"

	"github.com/parnurzeal/gorequest"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"
)

func GetBytes(url string) ([]byte, error) {
	log.Printf("GetBytes: " + url)
	resp, bytes, errs := gorequest.New().
		Get(url).
		Set("User-Agent", UserAgent).
		RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
			log.Printf("redirect: %s", req.URL.String())
			return nil
		}).
		EndBytes()
	if errs != nil {
		return nil, fmt.Errorf("get bytes: %s: %v", url, errs)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("get bytes: %s: %d %s", url, resp.StatusCode, resp.Status)
	}
	return bytes, nil
}
