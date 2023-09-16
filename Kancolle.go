package mediaunlocktest

import (
	"net/http"
)

func Kancolle(c http.Client) Result {
	resp, err := GET_Dalvik(c, "http://203.104.209.7/kcscontents/news/")
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return Result{Status: StatusOK}
	case 403:
		return Result{Status: StatusNo}
	}
	return Result{Status: StatusUnexpected}
}
