package chttp

import (
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
