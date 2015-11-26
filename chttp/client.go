package chttp

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

func HttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				deadline := time.Now().Add(timeout)
				c, err := net.DialTimeout(network, addr, timeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
			DisableKeepAlives:     false,
			ResponseHeaderTimeout: timeout,
			DisableCompression:    false,
		},
	}
}

func HttpProxyClient(timeout time.Duration, proxyURL string) (*http.Client, error) {
	proxy, err := url.Parse(proxyURL)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	proxyClient := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(timeout)
				c, err := net.DialTimeout(netw, addr, timeout)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
			DisableKeepAlives:     false,
			ResponseHeaderTimeout: timeout,
			DisableCompression:    false,
			Proxy:                 http.ProxyURL(proxy),
		},
	}
	return proxyClient, nil
}

func HttpGet(c *http.Client, url, referer string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.152 Safari/537.36")
	if len(referer) > 0 {
		req.Header.Set("Referer", referer)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Println("Client.Do(req) error: ", err)
		return ""
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println("ioutil ReadAll error: ", err)
		return ""
	}
	return string(body)
}
