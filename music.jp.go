package main

import (
	"io"
	"net/http"
)

func MusicJP(c http.Client) Result {
	resp, err := GET(c, "https://overseaauth.music-book.jp/globalIpcheck.js")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	if string(b) == "" {
		return Result{Success: false}
	}
	return Result{Success: true}
}
