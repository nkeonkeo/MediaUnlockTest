package main

import "net/http"

// World Flipper Japan
func WFJP(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://api.worldflipper.jp", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Dalvik)

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
