package main

import "net/http"

func PrettyDerbyJP(c http.Client) Result {
	resp, err := GET_Dalvik(c, "https://api-umamusume.cygames.jp/")
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
