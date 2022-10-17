package main

import "net/http"

// Princess Connect Re:Dive Japan
func PCRJP(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://api-priconne-redive.cygames.jp/", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("User-Agent", UA_Dalvik)

	resp, err := c.Do(req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	switch resp.StatusCode {
	case 404:
		return Result{Success: true}
	case 403:
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}
