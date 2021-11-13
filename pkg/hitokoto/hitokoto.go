package hitokoto

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
)

type Hitokoto struct {
	BaseURL   string
	SpanLimit time.Duration

	// sent query
	request chan struct{}
	// got error or	*Sentence
	response  chan interface{}
	lastFetch time.Time
}

var (
	Default = New()
)

func SayHi() {
	Default.request <- struct{}{}
}

func GetHitokoto() (*Sentence, error) {
	res := <-Default.response
	if err, ok := res.(error); ok {
		return nil, err
	}
	return res.(*Sentence), nil
}

func New() *Hitokoto {
	hitoko := &Hitokoto{
		BaseURL:   "https://v1.hitokoto.cn",
		SpanLimit: 300 * time.Millisecond,
		request:   make(chan struct{}),
		response:  make(chan interface{}),
	}
	go func() {
		for range hitoko.request {
			span := time.Since(hitoko.lastFetch)
			if span < hitoko.SpanLimit {
				time.Sleep(hitoko.SpanLimit - span)
			}
			hi, err := hitoko.Say()
			if err != nil {
				hitoko.response <- err
			} else {
				hitoko.response <- hi
			}
			hitoko.lastFetch = time.Now()
		}
	}()
	return hitoko
}

func (h *Hitokoto) fetch(param ...string) ([]byte, error) {
	uri := h.BaseURL
	if param != nil {
		uri += fmt.Sprintf("?%s", param[0])
	}
	resp, body, errs := gorequest.New().
		Timeout(3 * time.Second).
		Get(uri).EndBytes()
	if errs != nil {
		return nil, errors.Wrap(errs[0], "request")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response: %d %s", resp.StatusCode, resp.Status)
	}
	return body, nil
}

func (h *Hitokoto) Say() (*Sentence, error) {
	res, err := h.fetch()
	if err != nil {
		return nil, errors.Wrap(err, "fetch")
	}
	out := &Sentence{}
	err = json.Unmarshal(res, out)
	if err != nil {
		return nil, errors.Wrap(err, "sentence")
	}
	return out, nil
}
