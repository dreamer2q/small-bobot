package zhaosheng

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"strings"
)

const (
	baseUrl = "http://zhaosheng0.hdu.edu.cn/Template/Default/search.asp?Action=OK"
)

func GetQueryResult(ksh, name string) (string, error) {
	req := gorequest.New().Post(baseUrl).
		Type("form").
		Send(map[string]string{
			"txtKSH":  ksh,
			"txtname": name,
		})
	_, bodyBytes, errs := req.EndBytes()
	if errs != nil {
		return "", fmt.Errorf("query: %v", errs)
	}
	reader := transform.NewReader(bytes.NewReader(bodyBytes), simplifiedchinese.GBK.NewDecoder())
	docs, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", fmt.Errorf("err parsing html: %v", err)
	}
	res := docs.Find("body").Text()
	res = strings.ReplaceAll(res, "\n", "")
	return strings.TrimSpace(res), nil
}
