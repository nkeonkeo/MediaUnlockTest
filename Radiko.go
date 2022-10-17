package main

import (
	"io"
	"net/http"
	"strings"
)

func Radiko(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://radiko.jp/area?_=1625406539531", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	s := string(b)
	if strings.Contains(s, `classs="OUT"`) {
		return Result{Success: false}
	}
	if strings.Contains(s, "JAPAN") {
		return Result{Success: true}
	}
	return Result{Success: false}
}
