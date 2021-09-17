package baike

import (
	"errors"
	"fmt"
	"strings"

	"github.com/parnurzeal/gorequest"
)

const (
	baseURL   = "https://baike.baidu.com/item/"
	baseQuery = baseURL + "%s"
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36"
)

func getBaike(word string) ([]byte, error) {
	req := gorequest.
		New().
		Set("User-Agent", userAgent).
		Get(fmt.Sprintf(baseQuery, word)).
		RedirectPolicy(func(req gorequest.Request, via []gorequest.Request) error {
			if strings.Contains(req.URL.Path, "error.html") {
				return errors.New("keyword not found")
			}
			return nil
		})
	resp, bytes, errs := req.EndBytes()
	if errs != nil {
		return nil, fmt.Errorf("gorequest: %v", errs)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("gorequest: invalid http status code: %d %s",
			resp.StatusCode, resp.Status)
	}
	return bytes, nil
}
