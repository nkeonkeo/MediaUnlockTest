package main

import "net/http"

func Karaoke(c http.Client) Result {
	req, err := http.NewRequest("GET", "http://cds1.clubdam.com/vhls-cds1/site/xbox/sample_1.mp4.m3u8", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Browser)

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
