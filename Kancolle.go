package mediaunlocktest

import (
	"net/http"
)

func Kancolle(c http.Client) Result {
	resp, err := GET_Dalvik(c, "http://203.104.209.7/kcscontents/news/")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return Result{Success: true}
	case 403:
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}
