package main

import (
	"io"
	"net/http"
	"strings"
)

func FOD(c http.Client) Result {
	resp, err := GET(c, "https://geocontrol1.stream.ne.jp/fod-geo/check.xml?time=1624504256")
	if err != nil {
		return Result{Success: false, Err: err}
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	s := string(b)
	if strings.Contains(s, "true") {
		return Result{Success: true}
	}
	if strings.Contains(s, "false") {
		return Result{Success: true}
	}
	return Result{Success: false, Info: "unknown"}
}
