package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	UA_Browser = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36"
	UA_Dalvik  = "Dalvik/2.1.0 (Linux; U; Android 9; ALP-AL00 Build/HUAWEIALP-AL00)"
)

type Result struct {
	Success bool
	Region  string
	Info    string
	Err     error
}

var Dialer = &net.Dialer{
	Timeout: 5 * time.Second,
}
var IPv4Transport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		return Dialer.DialContext(ctx, "tcp4", addr)
	},
	// ForceAttemptHTTP2:     true,
	// TLSHandshakeTimeout: 30 * time.Second,
}

func UseLastResponse(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }

var Ipv4HttpClient = http.Client{
	Timeout:       5 * time.Second,
	CheckRedirect: UseLastResponse,
	Transport:     IPv4Transport,
}
var Ipv6Transport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		return Dialer.DialContext(ctx, "tcp6", addr)
	},
	// ForceAttemptHTTP2:     true,
	// TLSHandshakeTimeout: 10 * time.Second,
}
var Ipv6HttpClient = http.Client{
	Timeout:       5 * time.Second,
	CheckRedirect: UseLastResponse,
	Transport:     Ipv6Transport,
}

func GET(c http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("user-agent", UA_Browser)
	return cdo(c, req)
}

func GET_Dalvik(c http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("user-agent", UA_Dalvik)
	return cdo(c, req)
}

var ErrNetwork = errors.New("network error")

func cdo(c http.Client, req *http.Request) (resp *http.Response, err error) {
	for i := 0; i < 5; i++ {
		if resp, err = c.Do(req); err == nil {
			return resp, nil
		}
	}
	return nil, ErrNetwork
}
func PostJson(c http.Client, url string, data string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("user-agent", UA_Browser)

	return cdo(c, req)
}

func PostForm(c http.Client, url string, data string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("user-agent", UA_Browser)

	return cdo(c, req)
}
