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
	return string(body)
}

func GetSecondHtml(url string) {
	c := NewCookies()
	c.getUrlRespHtml("http://q.stock.sohu.com/cn/150224/index.shtml")
	log.Println("Second html:", c.getUrlRespHtml("http://q.stock.sohu.com/cn/150224/index.shtml"))
}
