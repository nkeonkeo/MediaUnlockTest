package main

import (
	"io"
	"net/http"
	"strings"
)

func Niconico(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://www.nicovideo.jp/watch/so40278367", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Add("user-agent", UA_Browser)

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	if strings.Contains(string(b), "同じ地域") {
		return Result{Success: false}
	}
	return Result{Success: true}
}
