package mediaunlocktest

import "net/http"

// World Flipper Japan
func WFJP(c http.Client) Result {
	resp, err := GET_Dalvik(c, "https://api.worldflipper.jp/")
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
