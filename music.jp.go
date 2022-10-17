package main

import (
	"io"
	"net/http"
)

func MusicJP(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://overseaauth.music-book.jp/globalIpcheck.js", nil)
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
	if string(b) == "" {
		return Result{Success: false}
	}
	return Result{Success: true}
}
