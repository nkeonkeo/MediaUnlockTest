package main

import (
	"io"
	"net/http"
	"strings"
)

func MyVideo(c http.Client) Result {
	c.CheckRedirect = nil
	resp, err := GET(c, "https://www.myvideo.net.tw/login.do")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	if strings.Contains(string(b), "serviceAreaBlock") {
		return Result{Success: false}
	}
	return Result{Success: true}
}
