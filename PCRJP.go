package main

import "net/http"

// Princess Connect Re:Dive Japan
func PCRJP(c http.Client) Result {
	resp, err := GET_Dalvik(c, "https://api-priconne-redive.cygames.jp/")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 404:
		return Result{Success: true}
	case 403:
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}
