package main

import (
	"io"
	"net/http"
	"strings"
)

func HuluJP(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://id.hulu.jp", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)
	c.CheckRedirect = nil

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	// log.Println(string(b))
	if strings.Contains(string(b), "restrict") {
		return Result{Success: false}
	}
	return Result{Success: true}
}
