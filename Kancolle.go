package main

import (
	"net/http"
)

func Kancolle(c http.Client) Result {
	req, err := http.NewRequest("GET", "http://203.104.209.7/kcscontents/news/", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("user-agent", UA_Dalvik)

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	switch resp.StatusCode {
	case 200:
		return Result{Success: true}
	case 403:
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}
