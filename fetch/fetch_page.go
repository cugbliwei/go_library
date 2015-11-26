package fetch

import (
	"go_library/chttp"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func FetchPage(url string) *goquery.Document {
	client := chttp.HttpClient(time.Second * 10)
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
