package ghttp

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type Cookies struct {
	cookies   []*http.Cookie
	cookieJar *cookiejar.Jar
}

func NewCookies() *Cookies {
	c := &Cookies{
		cookies: nil,
	}
	c.cookieJar, _ = cookiejar.New(nil)
	return c
}

func (c *Cookies) getUrlRespHtml(url string) string {
	client := &http.Client{
		CheckRedirect: nil,
		Jar:           c.cookieJar,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("NewRequest error :", err)
		return ""
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Host", "zhixing.court.gov.cn")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "http://zhixing.court.gov.cn/search/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36")
	//req.Header.Set("Cookie", "__jsluid=719e75faa43f0125b63dfede3c27f8d2; __jsl_clearance=1448802894.93|0|HUFkBjKoPW%2FVR8Z7M%2FFRg%2FGNbYE%3D")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do  error :", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("ReadAll  error :", err)
		return ""
	}
	c.cookies = c.cookieJar.Cookies(req.URL)
	log.Println("cookies:", c.cookies)
	return string(body)
}

func GetSecondHtml(url string) string {
	c := NewCookies()
	log.Println(c.getUrlRespHtml("http://zhixing.court.gov.cn/search"))
	return c.getUrlRespHtml(url)
}
