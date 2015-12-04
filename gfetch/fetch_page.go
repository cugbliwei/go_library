package gfetch

import (
	"go_library/ghttp"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//抓取一个网页，反正goquery.Document方便实时解析
func FetchPage(url string) *goquery.Document {
	client := ghttp.HttpClient(time.Second * 10)
	response, err := client.Get(url)
	if err != nil {
		log.Println("http get error: ", err)
		return nil
	}
	body, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		log.Println("goquery new Document error: ", err)
		return nil
	}
	return body
}
