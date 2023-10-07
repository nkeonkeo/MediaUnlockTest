package mediaunlocktest

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	Version          = "1.0"
	StatusOK         = 1
	StatusNetworkErr = -1
	StatusErr        = -2
	StatusRestricted = 2
	StatusNo         = 3
	StatusBanned     = 4
	StatusFailed     = 5
	StatusUnexpected = 6
)

type Result struct {
	Status int
	Region string
	Info   string
	Err    error
}

var (
	UA_Browser = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"
	UA_Dalvik  = "Dalvik/2.1.0 (Linux; U; Android 9; ALP-AL00 Build/HUAWEIALP-AL00)"
)

var Dialer = &net.Dialer{
	Timeout:   5 * time.Second,
	KeepAlive: 30 * time.Second,
}
var ipv4Transport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		return Dialer.DialContext(ctx, "tcp4", addr)
	},
	// ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   5 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	TLSClientConfig:       tlsConfig,
}

func UseLastResponse(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }

var defaultCipherSuites = []uint16{0xc02f, 0xc030, 0xc02b, 0xc02c, 0xcca8, 0xcca9, 0xc013, 0xc009, 0xc014, 0xc00a, 0x009c, 0x009d, 0x002f, 0x0035, 0xc012, 0x000a}

var Ipv4HttpClient = http.Client{
	Timeout:       10 * time.Second,
	CheckRedirect: UseLastResponse,
	Transport:     ipv4Transport,
}
var ipv6Transport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		return Dialer.DialContext(ctx, "tcp6", addr)
	},
	// ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   5 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	TLSClientConfig:       tlsConfig,
}
var Ipv6HttpClient = http.Client{
	Timeout:       10 * time.Second,
	CheckRedirect: UseLastResponse,
	Transport:     ipv6Transport,
}
var AutoHttpClient = http.Client{
	Timeout:       10 * time.Second,
	CheckRedirect: UseLastResponse,
	Transport: &http.Transport{
		Proxy:       http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
		// ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	},
}
var tlsConfig = &tls.Config{
	CipherSuites: append(defaultCipherSuites[8:], defaultCipherSuites[:8]...),
}

type H [2]string

func GET(c http.Client, url string, headers ...H) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("user-agent", UA_Browser)
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// req.Header.Set("accept-encoding", "gzip, deflate, br")
	// req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("dnt", "1")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	for _, h := range headers {
		req.Header.Set(h[0], h[1])
	}
	// return c.Do(req)
	return cdo(c, req)
}

func GET_Dalvik(c http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", UA_Dalvik)
	return cdo(c, req)
}

var ErrNetwork = errors.New("network error")

func cdo(c http.Client, req *http.Request) (resp *http.Response, err error) {
	// resp, err = c.Do(req)
	// if err != nil {
	// 	err = ErrNetwork
	// }
	// return
	deadline := time.Now().Add(30 * time.Second)
	for i := 0; i < 3; i++ {
		if resp, err = c.Do(req); err == nil {
			return resp, nil
		}
		if strings.Contains(err.Error(), "no such host") {
			break
		}
		if strings.Contains(err.Error(), "timeout") || time.Now().After(deadline) {
			break
		}
	}
	// log.Println(err)
	return nil, err
}
func PostJson(c http.Client, url string, data string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("user-agent", UA_Browser)
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// req.Header.Set("accept-encoding", "gzip, deflate, br")
	// req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("dnt", "1")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")

	return cdo(c, req)
}

func PostForm(c http.Client, url string, data string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("user-agent", UA_Browser)
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// req.Header.Set("accept-encoding", "gzip, deflate, br")
	// req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("dnt", "1")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")

	return cdo(c, req)
}
