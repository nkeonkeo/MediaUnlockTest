package main

import (
	"io"
	"net/http"
	"strings"
)

func HuluJP(c http.Client) Result {
	resp, err := GET(c, "https://id.hulu.jp")
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
