package main

import (
	"net/http"
	"strings"
)

func Hotstar(c http.Client) Result {
	r, err := http.NewRequest("GET", "https://api.hotstar.com/o/v1/page/1557?offset=0&size=20&tao=0&tas=20", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	r.Header.Set("User-Agent", UA_Browser)

	resp, err := c.Do(r)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	switch resp.StatusCode {
	case 475:
		return Result{Success: false}
	case 401:
		r, err := http.NewRequest("GET", "https://www.hotstar.com", nil)
		if err != nil {
			return Result{Success: false, Err: err}
		}
		r.Header.Set("User-Agent", UA_Browser)
		resp, err := c.Do(r)
		if err != nil {
			return Result{Success: false, Err: err}
		}
		if resp.StatusCode == 301 {
			return Result{Success: false}
		}
		u := resp.Header.Get("Location")
		if u == "" {
			return Result{Success: false}
		}
		t := strings.SplitN(u, "/", 4)
		if len(t) < 4 {
			return Result{Success: false}
		}
		return Result{Success: true, Region: t[3]}
	}
	return Result{Success: false, Info: "Failed"}
}
