package main

import (
	"io"
	"net/http"
	"strings"
)

func MyVideo(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://www.myvideo.net.tw/login.do", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	c.CheckRedirect = nil

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	if strings.Contains(string(b), "serviceAreaBlock") {
		return Result{Success: false}
	}
	return Result{Success: true}
}
