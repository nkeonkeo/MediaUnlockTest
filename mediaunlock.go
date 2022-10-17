package main

import (
	"context"
	"net"
	"net/http"
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
	Timeout:   5 * time.Second,
	KeepAlive: 30 * time.Second,
}
var IPv4Transport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
		return Dialer.DialContext(ctx, "tcp4", addr)
	},
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   5 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
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
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   5 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
var Ipv6HttpClient = http.Client{
	Timeout:       5 * time.Second,
	CheckRedirect: UseLastResponse,
	Transport:     Ipv6Transport,
}
